package config

type Config struct {
	ProxmoxHost     string `env:"PROXMOX_HOST"`
	ProxmoxApiToken string `env:"PROXMOX_API_TOKEN"`
}

func init() {

}
