package k8s_test

import (
	. "github.com/cloudfoundry/uaa/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"path/filepath"
)

var _ = Describe("Secrets", func() {

	var templates []string

	It("Renders with SMTP credentials", func() {
		templates = []string{
			pathToFile(filepath.Join("values", "_values.yml")),
			pathToFile(filepath.Join("secrets", "smtp_credentials.yml")),
		}

		renderingContext := NewRenderingContext(templates...).WithData(
			map[string]string{
				"smtp.user":     "my smtp username",
				"smtp.password": "my smtp password",
				"smtp.host":     "my smtp host",
				"smtp.port":     "my smtp port",
				"smtp.starttls": "my smtp starttls",
			})

		smtp_secrets := `smtp:
  user: my smtp username
  password: my smtp password
  host: my smtp host
  port: my smtp port
  starttls: my smtp starttls
`

		Expect(renderingContext).To(
			ProduceYAML(RepresentingASecret().
				WithName("uaa-smtp-credentials").
				WithStringData("smtp_credentials.yml", smtp_secrets)),
		)
	})

	It("Renders with Database credentials", func() {
		templates = []string{
			pathToFile(filepath.Join("values", "_values.yml")),
			pathToFile(filepath.Join("secrets", "database_credentials.yml")),
		}

		renderingContext := NewRenderingContext(templates...).WithData(
			map[string]string{
				"database.username": "my database username",
				"database.password": "my database password",
			})

		database_credentials := `database:
  username: my database username
  password: my database password
`

		Expect(renderingContext).To(
			ProduceYAML(RepresentingASecret().
				WithName("uaa-database-credentials").
				WithStringData("database_credentials.yml", database_credentials)),
		)
	})

	It("Renders with Different Database credentials", func() {
		templates = []string{
			pathToFile(filepath.Join("values", "_values.yml")),
			pathToFile(filepath.Join("secrets", "database_credentials.yml")),
		}

		renderingContext := NewRenderingContext(templates...).WithData(
			map[string]string{
				"database.username": "my other database username",
				"database.password": "my other database password",
			})

		database_credentials := `database:
  username: my other database username
  password: my other database password
`

		Expect(renderingContext).To(
			ProduceYAML(RepresentingASecret().
				WithName("uaa-database-credentials").
				WithStringData("database_credentials.yml", database_credentials)),
		)
	})

})
