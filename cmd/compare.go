package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/nsf/jsondiff"
	"github.com/spf13/cobra"
	"os"
)

var compareCmd = &cobra.Command{
	Use:   "compare [file1] [file2]",
	Short: "Compare two workflow backup files",
	Long:  `Compare two workflow backup JSON files and show the differences.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		runCompare(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(compareCmd)
}

func runCompare(file1, file2 string) {
	fmt.Println("🔍 Comparing workflow backups...")
	fmt.Println()

	data1, err := os.ReadFile(file1)
	if err != nil {
		color.Red("❌ Error reading file 1:")
		fmt.Printf("   %v\n", err)
		return
	}

	data2, err := os.ReadFile(file2)
	if err != nil {
		color.Red("❌ Error reading file 2:")
		fmt.Printf("   %v\n", err)
		return
	}

	// Parse JSON to get workflow names
	var workflow1, workflow2 map[string]interface{}
	json.Unmarshal(data1, &workflow1)
	json.Unmarshal(data2, &workflow2)

	name1 := getWorkflowName(workflow1)
	name2 := getWorkflowName(workflow2)

	// Display file info
	color.Cyan("File 1: %s", file1)
	fmt.Printf("  Workflow: %s\n\n", name1)

	color.Cyan("File 2: %s", file2)
	fmt.Printf("  Workflow: %s\n\n", name2)

	// Compare JSON
	opts := jsondiff.DefaultConsoleOptions()
	result, diff := jsondiff.Compare(data1, data2, &opts)

	// Display results
	fmt.Println("═══════════════════════════════════════")
	fmt.Println("Comparison Result:")
	fmt.Println("═══════════════════════════════════════")

	switch result {
	case jsondiff.FullMatch:
		color.Green("✅ IDENTICAL - Both workflows are exactly the same!")

	case jsondiff.SupersetMatch:
		color.Yellow("⚠️  SUPERSET - File 1 contains all of File 2's content plus more")
		fmt.Println()
		fmt.Println("Differences:")
		fmt.Println(diff)

	case jsondiff.NoMatch:
		color.Red("❌ DIFFERENT - Workflows have differences")
		fmt.Println()
		fmt.Println("Differences:")
		fmt.Println(diff)
		showDetailedDiff(workflow1, workflow2)

	default:
		color.Red("❌ Error comparing files: Invalid JSON")
	}

	fmt.Println()
}

func getWorkflowName(workflow map[string]interface{}) string {
	if name, ok := workflow["name"].(string); ok {
		return name
	}
	return "Unknown"
}

func showDetailedDiff(wf1, wf2 map[string]interface{}) {
	fmt.Println()
	color.Cyan("📊 Detailed Analysis:")
	fmt.Println()

	// Compare node counts
	nodes1, _ := wf1["nodes"].([]interface{})
	nodes2, _ := wf2["nodes"].([]interface{})

	if len(nodes1) != len(nodes2) {
		fmt.Printf("Nodes: ")
		if len(nodes1) > len(nodes2) {
			color.Green("%d", len(nodes1))
		} else {
			color.Red("%d", len(nodes1))
		}
		fmt.Printf(" → ")
		if len(nodes2) > len(nodes1) {
			color.Green("%d", len(nodes2))
		} else {
			color.Red("%d", len(nodes2))
		}
		fmt.Printf(" (Δ %+d)\n", len(nodes2)-len(nodes1))
	} else {
		fmt.Printf("Nodes: %d (unchanged)\n", len(nodes1))
	}

	// Compare active status
	active1, _ := wf1["active"].(bool)
	active2, _ := wf2["active"].(bool)

	if active1 != active2 {
		fmt.Printf("Status: ")
		if active1 {
			color.Green("Active")
		} else {
			color.Yellow("Inactive")
		}
		fmt.Printf(" → ")
		if active2 {
			color.Green("Active")
		} else {
			color.Yellow("Inactive")
		}
		fmt.Println()
	}

	// Compare update times
	updated1, _ := wf1["updatedAt"].(string)
	updated2, _ := wf2["updatedAt"].(string)

	if updated1 != updated2 {
		fmt.Printf("Updated: %s → %s\n", updated1, updated2)
	}
}
