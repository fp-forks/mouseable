// Package def declares types for prevent circular dependencies.
package def

import (
	"github.com/wirekang/mouseable/internal/di"
)

func New() di.DefinitionManager {
	m := &manager{
		keyStringCmdNameMap: map[di.CommandKeyString]di.CommandName{},
		cmdDefMap:           map[di.CommandName]*commandDef{},
		dataDefMap:          map[di.DataName]*dataDef{},
	}
	m.insertCommand(
		"activate",
		"Activate Mouseable",
		di.WhenDeactivated,
		&di.Command{
			OnBegin: func(tool *di.CommandTool) {
				tool.Activate()
			},
		},
	)
	m.insertCommand(
		"deactivate",
		"Deactivate Mouseable",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"move-right",
		"Move cursor →",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"move-right-up",
		"Move cursor ↗",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"move-up",
		"Move cursor ↑",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"move-left-up",
		"Move cursor ↖",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"move-left",
		"Move cursor ←",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"move-left-down",
		"Move cursor ↙",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"move-down",
		"Move cursor ↓",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"move-right-down",
		"Move cursor ↘",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"sniper-mode",
		"Slow down to increase accuracy",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"sniper-mode-wheel",
		"Slow down to increase accuracy (MouseWheel)",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"click-left",
		"Click left mouse button",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"click-right",
		"Click right mouse button",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"click-middle",
		"Click middle mouse button",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"wheel-up",
		"MouseWheel ↑",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"wheel-down",
		"MouseWheel ↓",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"wheel-right",
		"MouseWheel →",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"wheel-left",
		"MouseWheel ←",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"teleport-forward",
		"Teleport cursor to the direction it is moving",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"teleport-right",
		"Teleport cursor →",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"teleport-right-up",
		"Teleport cursor ↗",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"teleport-up",
		"Teleport cursor ↑",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"teleport-left-up",
		"Teleport cursor ↖",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"teleport-left",
		"Teleport cursor ←",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"teleport-left-down",
		"Teleport cursor ↙",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"teleport-down",
		"Teleport cursor ↓",
		di.WhenActivated,
		&di.Command{},
	)
	m.insertCommand(
		"teleport-right-down",
		"Teleport cursor ↘",
		di.WhenActivated,
		&di.Command{},
	)

	m.insertData("key-timeout", "Key press timeout for continuous input in ms", di.Int, 200)
	m.insertData("cursor-acceleration-x", "Cursor horizontal acceleration", di.Float, 2.8)
	m.insertData("cursor-acceleration-y", "Cursor vertical acceleration", di.Float, 2.8)
	m.insertData("cursor-friction-x", "Cursor horizontal friction", di.Float, 2.5)
	m.insertData("cursor-friction-y", "Cursor vertical friction", di.Float, 2.5)
	m.insertData("wheel-acceleration-x", "MouseWheel horizontal acceleration", di.Int, 5)
	m.insertData("wheel-acceleration-y", "MouseWheel vertical acceleration", di.Int, 5)
	m.insertData("wheel-friction-x", "MouseWheel horizontal friction", di.Int, 4)
	m.insertData("wheel-friction-y", "MouseWheel vertical friction", di.Int, 4)
	m.insertData("cursor-sniper-speed-x", "Sniper mode horizontal speed", di.Int, 1)
	m.insertData("cursor-sniper-speed-y", "Sniper mode vertical speed", di.Int, 1)
	m.insertData("wheel-sniper-speed-x", "Sniper mode horizontal speed (MouseWheel)", di.Int, 1)
	m.insertData("wheel-sniper-speed-y", "Sniper mode vertical speed (MouseWheel)", di.Int, 1)
	m.insertData("teleport-distance-f", "TeleportForward distance", di.Int, 300)
	m.insertData("teleport-distance-x", "Teleport horizontal distance", di.Int, 300)
	m.insertData("teleport-distance-y", "Teleport vertical distance", di.Int, 300)
	m.insertData("show-overlay", "Show overlay when Mouseable activated", di.Bool, true)

	return m
}
