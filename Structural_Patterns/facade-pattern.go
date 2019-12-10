package Structural_Patterns

type Music struct {
	Name string
}
func (m *Music) GetMusic() string {
	return m.Name
}



//Video struct
type Video struct {
	Id int64
}
func (v *Video) GetVideoId() int64 {
	return v.Id
}

//count struct
type Count struct {
	Comment int64
	Praise 	int64
	Collect int64
}
func (c *Count) GetComment() int64 {
	return c.Comment
}




//外观结构Facade
type Facade struct {
	music Music
	video Video
	count Count
}
//对外访问接口
func (f *Facade) PrintServerInfo() {
	f.music.GetMusic()
	f.video.GetVideoId()
	f.count.GetComment()
}
func NewFacade(music Music, count Count, video Video) *Facade {
	return &Facade{
		music: music,
		video: video,
		count: count,
	}
}