package library

import "testing"

func TestOpts(t *testing.T) {
	mm := NewMusicManager()

	if mm == nil {
		t.Error("NewMusicManager Failed")
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManager failed, not empty.")
	}

	m0 := &MusicEntry{
		"1", "My Heart Will Go On", "Celion Diaon", Pop,
		"http://music.baidu.com/data/music/file?link=http://zhangmenshiting.baidu.com/data2/music/124574160/124574160.mp3?xcode=49549704596d112696a2d3f13bf86a03",
		MP3}

	mm.Add(m0)
	if mm.Len() != 1 {
		t.Error("MusicManager.Add Failed.")
	}

	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("MusicManager.Find() Failed")
	}
	if m.Id != m0.Id || m.Artist != m0.Artist ||
		m.Name != m0.Name || m.Source != m0.Source ||
		m.Type != m0.Type || m.Genre != m0.Genre {
		t.Error("MusicManager.Find() Failed. Found item mismatch.")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusicManager.Get() Failed.", err)
	}

	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove() Failed.", err)
	}

}
