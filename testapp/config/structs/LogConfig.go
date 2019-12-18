package config

type LogConfiguration struct {
	FilePath			string	`json:"file_path"`
	LogFileName			string	`json:"log_file_name"`
	ErrorLogFileName	string	`json:"error_log_file_name"`
	DebugLogFileName	string	`json:"debug_log_file_name"`
	MaxSize				int		`json:"max_size"`
	MaxBackup			int		`json:"max_backup"`
	MaxAge				int		`json:"max_age"`
	Compress			bool	`json:"compress"`
	Level				string	`json:"level"`
}
