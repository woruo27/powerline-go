package main

import pwl "github.com/justjanne/powerline-go/powerline"

func segmentRoot(p *powerline) []pwl.Segment {
	var foreground, background uint8
	var content string
	if p.cfg.PrevError == 0 || p.cfg.StaticPromptIndicator {
		foreground = p.theme.CmdPassedFg
		background = p.theme.CmdPassedBg
		content    = "٩(>ω<)۶"
	} else {
		foreground = p.theme.CmdFailedFg
		background = p.theme.CmdFailedBg
		content    = "(º﹃º)"
	}

	return []pwl.Segment{{
		Name:       "root",
		// Content:    p.shell.RootIndicator,
		Content:    content,
		Foreground: foreground,
		Background: background,
	}}
}
