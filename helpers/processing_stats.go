// Copyright 2017 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package helpers

import (
	"io"
	"strconv"
	"sync/atomic"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
)

// ProcessingStats represents statistics about a site build.
type ProcessingStats struct {
	Name string

	Pages           uint64
	PaginatorPages  uint64
	Static          uint64
	ProcessedImages uint64
	Files           uint64
	Aliases         uint64
	Cleaned         uint64
}

type processingStatsTitleVal struct {
	name string
	val  uint64
}

func (s *ProcessingStats) toVals() []processingStatsTitleVal {
	return []processingStatsTitleVal{
		{"Pages", s.Pages},
		{"Paginator pages", s.PaginatorPages},
		{"Non-page files", s.Files},
		{"Static files", s.Static},
		{"Processed images", s.ProcessedImages},
		{"Aliases", s.Aliases},
		{"Cleaned", s.Cleaned},
	}
}

// NewProcessingStats returns a new ProcessingStats instance.
func NewProcessingStats(name string) *ProcessingStats {
	return &ProcessingStats{Name: name}
}

// Incr increments a given counter.
func (s *ProcessingStats) Incr(counter *uint64) {
	atomic.AddUint64(counter, 1)
}

// Add adds an amount to a given counter.
func (s *ProcessingStats) Add(counter *uint64, amount int) {
	atomic.AddUint64(counter, uint64(amount))
}

// ProcessingStatsTable writes a table-formatted representation of stats to w.
func ProcessingStatsTable(w io.Writer, stats ...*ProcessingStats) {
	names := make([]string, len(stats)+1)

	var data [][]string

	for i := range stats {
		stat := stats[i]
		names[i+1] = stat.Name

		titleVals := stat.toVals()

		if i == 0 {
			data = make([][]string, len(titleVals))
		}

		for j, tv := range titleVals {
			if i == 0 {
				data[j] = []string{tv.name, strconv.Itoa(int(tv.val))}
			} else {
				data[j] = append(data[j], strconv.Itoa(int(tv.val)))
			}
		}
	}

	table := tablewriter.NewTable(
		w,
		tablewriter.WithRenderer(renderer.NewBlueprint(tw.Rendition{
			Borders: tw.BorderNone,
			Symbols: tw.NewSymbols(tw.StyleLight),
			Settings: tw.Settings{
				Separators: tw.Separators{BetweenRows: tw.Off},
				Lines:      tw.Lines{ShowFooterLine: tw.On},
			},
		})),
		tablewriter.WithConfig(
			tablewriter.Config{
				MaxWidth: 70,
				Row:      tw.CellConfig{Alignment: tw.CellAlignment{Global: tw.AlignRight, PerColumn: []tw.Align{tw.AlignLeft}}},
			}),
	)

	table.Bulk(data)
	table.Header(names)
	table.Render()
}
