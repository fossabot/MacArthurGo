package cqcode

import "time"

func At(qq int64) string {
	data := map[string]any{
		"qq": qq,
	}
	cq := CQCode{Type: "at", Data: data}
	return cq.toString()
}

func Reply(msgId int64, text string, qq int64, seq int64) string {
	data := map[string]any{
		"id":   msgId,
		"text": text,
		"qq":   qq,
		"time": time.Now().Unix(),
		"seq":  seq,
	}
	cq := CQCode{Type: "reply", Data: data}
	return cq.toString()
}

func Poke(qq int64) string {
	data := map[string]any{
		"qq": qq,
	}
	cq := CQCode{Type: "poke", Data: data}
	return cq.toString()
}

func Music(urlType string, id int64) string {
	data := map[string]any{
		"type": urlType,
		"id":   id,
	}
	cq := CQCode{Type: "music", Data: data}
	return cq.toString()
}

func Image(file string) string {
	data := map[string]any{
		"file": file,
	}
	cq := CQCode{Type: "image", Data: data}
	return cq.toString()
}
