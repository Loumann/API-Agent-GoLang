package repos

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Repository struct {
	db *sqlx.DB
}

var db *sqlx.DB
var dbDriveName = "postgres"

func DbConnection(config *config.Config) (*sqlx.DB, error) {
	return sqlx.Connect(dbDriveName,
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.Username, config.Password, config.Dbname, config.SSLmode))
}
func GetRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetUsers() ([]models.Agent, error) {
	var agents []models.Agent

	row, err := r.db.Query("SELECT id, agentname, status FROM agents")
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		var agent models.Agent
		if err := row.Scan(&agent.ID, &agent.AgentName, &agent.Status); err != nil {
			return nil, err
		}
		agents = append(agents, agent)

	}
	return agents, nil
}
func (r *Repository) UpdateAgent(agent models.Agent) (bool, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM agents WHERE id=$1", agent.ID).Scan(&count)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}

	query := `UPDATE agents SET agentname=$1, status=$2 WHERE id=$3`
	_, err = r.db.Exec(query, agent.AgentName, agent.Status, agent.ID)
	return true, err
}
func (r *Repository) DeleteUser(id int) error {
	_, err := r.db.Exec("DELETE FROM agents WHERE id=$1", id)
	if err != nil {
		log.Println(err)
	}

	return err
}
func (r *Repository) Create(agentname, status string) error {
	var id int
	err := r.db.QueryRow(`insert into "agents" ("agentname", "status") values ($1,$2) RETURNING id`,
		agentname, status).Scan(&id)
	if err != nil {
	}
	return nil
}
