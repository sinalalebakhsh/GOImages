package config




// The Configuration interface defines methods for retrieving configuration settings, 
// with support for obtaining string, int, float64, and bool values. 
// There is also a set of methods that allow a default value to
// be supplied. The configuration data will allow nested configuration sections, 
// which can be obtained using the GetSection method.
type Configuration interface {
	GetString(name string) (configValue string, found bool)
	GetInt(name string) (configValue int, found bool)
	GetBool(name string) (configValue bool, found bool)
	GetFloat(name string) (configValue float64, found bool)
	GetStringDefault(name, defVal string) (configValue string)
	GetIntDefault(name string, defVal int) (configValue int)
	GetBoolDefault(name string, defVal bool) (configValue bool)
	GetFloatDefault(name string, defVal float64) (configValue float64)
	GetSection(sectionName string) (section Configuration, found bool)
}
