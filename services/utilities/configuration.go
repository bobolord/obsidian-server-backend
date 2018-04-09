package utilities

type ConfigStruct struct {
	DbmsConfig   DbmsConfigStruct   `yaml:"dbms_Config,omitempty"`
	AppConfig    AppConfigStruct    `yaml:"app_Config,omitempty"`
	MailerConfig MailerConfigStruct `yaml:"mailer_Config,omitempty"`
}

type DbmsConfigStruct struct {
	Dbms     string `yaml:"dbms,omitempty"`
	Host     string `yaml:"host,omitempty"`
	Port     string `yaml:"port,omitempty"`
	Database string `yaml:"database,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
}

type AppConfigStruct struct {
	AllowedOrigins  string `yaml:"allowed_origins,omitempty"`
	Port            string `yaml:"port,omitempty"`
	Domain          string `yaml:"domain,omitempty"`
	CsrfTokenExpiry int    `yaml:"csrfTokenExpiry,omitempty"`
}

type MailerConfigStruct struct {
	SenderEmail string `yaml:"sender_email,omitempty"`
	Password    string `yaml:"password,omitempty"`
}

var Config ConfigStruct
