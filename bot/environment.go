package bot

type Environment struct {
	Config            *Config
	Storage           Storage
	Resources         *Resources
	TempData          TempStorage
	CallBackProcessor CallBackProcessor
	MessageProcessor  MessageProcessor
	ButtonMaker       ButtonMaker
}
