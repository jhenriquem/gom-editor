package screen

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/Gom/config"
)

var Screen, err = tcell.NewScreen()

func ScreenInitializer() {
	if err != nil {
		log.Fatalf("Erro ao iniciar a tela: %v", err)
	}

	if err := Screen.Init(); err != nil {
		log.Fatalf("Erro ao inicializar a tela: %v", err)
	}

	// Definer o estilo da tela
	stScreen := tcell.StyleDefault.Background(config.ColorBg).Foreground(tcell.ColorReset)
	Screen.SetStyle(stScreen)
	Screen.Clear()
	Screen.Show()
}
