package cmd

import (
	"fmt"
	"godevops/internal/system"
	"time"

	"godevops/internal/ui"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Interactive system dashboard",
	Run: func(cmd *cobra.Command, args []string) {

		app := tview.NewApplication()

		// painel de sistema
		systemBox := tview.NewTextView().
			SetDynamicColors(true).
			SetChangedFunc(func() {
				app.Draw()
			})
		systemBox.SetBorder(true).SetTitle("FULLDEVSTACKS SYSTEM METRICS")

		// painel de rede
		networkBox := tview.NewTextView().
			SetDynamicColors(true).
			SetChangedFunc(func() {
				app.Draw()
			})
		networkBox.SetBorder(true).SetTitle("NETWORK")

		// painel de processos
		processBox := tview.NewTextView().
			SetDynamicColors(true).
			SetChangedFunc(func() {
				app.Draw()
			})
		processBox.SetBorder(true).SetTitle("TOP PROCESSES")

		// layout vertical
		layout := tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(systemBox, 0, 2, false).
			AddItem(networkBox, 0, 1, false).
			AddItem(processBox, 0, 2, false)

		go func() {
			for {

				cpu := system.GetCPUUsage()
				mem := system.GetMemoryUsage()
				disk := system.GetDiskUsage()
				proc := system.GetProcessCount()
				up, down := system.GetNetworkSpeed()

				upMB := system.BytesToMB(up)
				downMB := system.BytesToMB(down)

				top := system.GetTopProcesses(5)

				app.QueueUpdateDraw(func() {

					systemBox.Clear()
					networkBox.Clear()
					processBox.Clear()

					// SYSTEM PANEL
					fmt.Fprintf(systemBox,
						"\nCPU: %s %.2f%%"+
							"\nRAM: %s %.2f%%"+
							"\nDISK: %s %.2f%%"+
							"\nPROC: %d",
						ui.CreateBar(cpu, 20), cpu,
						ui.CreateBar(mem, 20), mem,
						ui.CreateBar(disk, 20), disk,
						proc)

					// NETWORK PANEL
					fmt.Fprintf(networkBox,
						"\nDOWNLOAD ↓ %.2fMB/s"+
							"\nUPLOAD ↑ %.2fMB/s\n",
						downMB, upMB)

					// PROCESS HEADER
					fmt.Fprintf(processBox, "%-20s %s\n", "PROCESS", "CPU")
					fmt.Fprintf(processBox, "────────────────────────\n")

					for _, p := range top {
						fmt.Fprintf(processBox, "%-20s %.2f%%\n", p.Name, p.CPU)
					}
				})

				time.Sleep(1 * time.Second)
			}
		}()

		// CAPTURA TECLADO
		app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Rune() {
			case 'q': // sair
				app.Stop()
				return nil
			case 'r': // redraw
				app.Draw()
				return nil
			}
			return event
		})

		// inicia a aplicação
		if err := app.SetRoot(layout, true).Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dashboardCmd)
}
