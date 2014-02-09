package main

import (
	"dlib/dbus"
	"fmt"
)

func (theme *Theme) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{
		"com.deepin.daemon.Grub2",
		fmt.Sprintf("/com/deepin/daemon/Grub2/Theme%d", theme.id),
		"com.deepin.daemon.Grub2.Theme",
	}
}

func (theme *Theme) OnPropertiesChanged(name string, oldv interface{}) {
	defer func() {
		if err := recover(); err != nil {
			logError("%v", err)
		}
	}()
	switch name {
	case "Background":
		theme.setBackground(theme.Background)
	case "ItemColor":
		theme.setItemColor(theme.ItemColor)
	case "SelectedItemColor":
		theme.setSelectedItemColor(theme.SelectedItemColor)
	}
}

func (theme *Theme) Reset() error {
	tplJsonData, err := theme.getThemeTplJsonData()
	if err != nil {
		return err
	}

	theme.relBgFile = tplJsonData.DefaultTplValue.Background
	theme.makeAbsBgFile()
	theme.ItemColor = tplJsonData.DefaultTplValue.ItemColor
	theme.SelectedItemColor = tplJsonData.DefaultTplValue.SelectedItemColor

	theme.customTheme()
	return nil
}
