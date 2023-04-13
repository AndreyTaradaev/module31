package postgres

import (
	"context"
	"db/pkg/typetask"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func New(constr string) (*Storage, error) {
	db, err := pgxpool.New(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) NewTask(t typetask.Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO tasks (title, content,author_id,assigned_id,"Opened")
		VALUES ($1, $2, $3, $4,$5) RETURNING id;
		`,
		t.Title,
		t.Content,
		t.AuthtorId,
		t.AssignedId,
		t.Opened,
	).Scan(&id)
	return id, err
}

// Tasks возвращает список задач из БД.
func (s *Storage) Tasks(taskID, authorID, label_id int) ([]typetask.Task, error) {
	rows, err := s.db.Query(context.Background(), `
	SELECT 
		id,
		"Opened",
		"Closed",
		author_id,
		assigned_id,
		title,
		content
	FROM tasks
	WHERE
		($1 = 0 OR id = $1) AND
		($2 = 0 OR author_id = $2) and
		(id in (select distinct task_id from tasks_labels tl
			where  tl.label_id  = $3) or  $3 = 0 
		)
	ORDER BY id;
`,
		taskID,
		authorID,
		label_id,
	)
	if err != nil {
		return nil, err
	}
	var tasks []typetask.Task
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную

	for rows.Next() {
		var t typetask.Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthtorId,
			&t.AssignedId,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		tasks = append(tasks, t)
	}
	// ВАЖНО не забыть проверить rows.Err()
	return tasks, rows.Err()
}

func (s *Storage) DeleteTask(id int) error {
	_, err := s.db.Exec(context.Background(),
		`	
						delete from tasks 
					WHERE  id = $1 ;		
		`,
		id,
	)
	return err
}

// изменяем задачу
func (s *Storage) UpdateTask(t typetask.Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
	UPDATE tasks
SET   "Closed" = $1, author_id = $2, assigned_id = $3, title = $4, content = $5
WHERE  id = $6  RETURNING id;		
		`,
		t.Closed,
		t.AuthtorId,
		t.AssignedId,
		t.Title,
		t.Content,
		t.ID,
	).Scan(&id)
	return id, err
}
