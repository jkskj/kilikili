package service

import (
	"kilikili/model"
	"kilikili/util/e"
	"kilikili/util/serializer"
)

type InteractionService struct {
	Vid  uint  `form:"vid"json:"vid"`   //视频id
	Type int64 `form:"type"json:"type"` //互动类型
}

func (service *InteractionService) Interact(uid uint) serializer.Response {
	code := e.SUCCESS
	var interaction model.Interaction
	var video model.Video
	var user model.User
	count := 0
	model.DB.First(&video, service.Vid)
	model.DB.First(&user, uid)
	model.DB.Where("vid=?", service.Vid).Where("type=?", service.Type).Where("uid=?", uid).First(&interaction).Count(&count)
	if count == 1 {
		if service.Type != 2 {
			if interaction.Status != 0 {
				code = e.ErrorExistInteraction
			} else {
				interaction.Status = 1
				model.DB.Save(interaction)
				if service.Type == 1 {
					video.LikeNum++
				} else {
					video.CollectionNum++
				}
				model.DB.Save(video)
			}
		} else {
			if interaction.Status == 0 {
				interaction.Status = 1
				model.DB.Save(interaction)
				video.ForwardNum++
				model.DB.Save(video)
			}
		}
	} else {
		interaction.Uid = uid
		interaction.Vid = service.Vid
		interaction.Type = service.Type
		interaction.Video = video
		interaction.User = user
		model.DB.Save(&interaction)
		if service.Type == 1 {
			video.LikeNum++
		} else if service.Type == 0 {
			video.CollectionNum++
		} else {
			video.ForwardNum++
		}
		model.DB.Save(video)
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
func (service *InteractionService) Cancel(uid uint) serializer.Response {
	code := e.SUCCESS
	var interaction model.Interaction
	count := 0
	var video model.Video
	model.DB.First(&video, service.Vid)
	model.DB.Where("vid=?", service.Vid).Where("type=?", service.Type).Where("uid=?", uid).First(&interaction).Count(&count)
	if count != 1 || interaction.Status == 0 {
		code = e.ErrorNotExistInteraction
	} else {
		interaction.Status = 0
		model.DB.Save(interaction)
		if service.Type == 1 {
			video.LikeNum--
		} else if service.Type == 0 {
			video.CollectionNum--
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
