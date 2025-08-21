package service

import (
	"context"
	"whatsapp-api/internal/repository"
	"whatsapp-api/model"
	proto "whatsapp-api/model/pb"
	"whatsapp-api/util"
)

func (s *service) PushMessage(ctx context.Context, msg *model.MessageRequest) (*model.MessageResponse, error) {

	if msg.DeviceID == "" {
		return nil, ErrDeviceIDEmpty
	}
	if msg.To == "" {
		return nil, ErrToEmpty
	}
	if msg.Type == "" {
		return nil, ErrTypeEmpty
	}

	// Validasi image
	if msg.Image != nil {
		if msg.Image.Url == "" {
			return nil, ErrImageUrlEmpty
		}
		if msg.Image.Caption == "" {
			return nil, ErrImageCaptionEmpty
		}
		if msg.Image.Mimetype == "" {
			return nil, ErrImageMimetypeEmpty
		}
	}
	// Validasi video
	if msg.Video != nil {
		if msg.Video.Url == "" {
			return nil, ErrVideoUrlEmpty
		}
		if msg.Video.Caption == "" {
			return nil, ErrVideoCaptionEmpty
		}
		if msg.Video.Mimetype == "" {
			return nil, ErrVideoMimetypeEmpty
		}
	}
	// Validasi document
	if msg.Document != nil {
		if msg.Document.Url == "" {
			return nil, ErrDocumentUrlEmpty
		}
		if msg.Document.Filename == "" {
			return nil, ErrDocumentFilenameEmpty
		}
		if msg.Document.Mimetype == "" {
			return nil, ErrDocumentMimetypeEmpty
		}
		if msg.Document.Title == "" {
			return nil, ErrDocumentTitleEmpty
		}
	}
	// Validasi audio
	if msg.Audio != nil {
		if msg.Audio.Url == "" {
			return nil, ErrAudioUrlEmpty
		}
		if msg.Audio.MimeType == "" {
			return nil, ErrAudioMimetypeEmpty
		}
		// ptt boolean tidak perlu validasi required
	}
	// Validasi location
	if msg.Location != nil {
		if msg.Location.Latitude == 0 {
			return nil, ErrLocationLatitudeEmpty
		}
		if msg.Location.Longitude == 0 {
			return nil, ErrLocationLongitudeEmpty
		}
		if msg.Location.Name == "" {
			return nil, ErrLocationNameEmpty
		}
		if msg.Location.Address == "" {
			return nil, ErrLocationAddressEmpty
		}
	}

	device, err := s.repo.FindDeviceByID(ctx, msg.DeviceID)
	if err != nil {
		if err == repository.ErrDeviceNotFound {
			return nil, ErrDeviceNotFound
		}
		return nil, err
	}

	message, err := s.grpc.SendMessage(ctx, &proto.MessagePayload{
		SenderJid: device.SenderJID,
		To:        util.EnsureWhatsAppJID(msg.To),
		Type:      msg.Type,
		Text:      msg.Text,
		Image:     msg.Image,
		Video:     msg.Video,
		Document:  msg.Document,
		Audio:     msg.Audio,
		Location:  msg.Location,
	})
	if err != nil {
		return nil, err
	}

	return &model.MessageResponse{
		Message: model.Message{
			ID: message.Id,
		},
	}, nil
}

func (s *service) GetContacts(ctx context.Context, deviceID string) (*model.Contacts, error) {

	if deviceID == "" {
		return nil, ErrDeviceIDEmpty
	}

	// Cek device terlebih dahulu
	device, err := s.repo.FindDeviceByID(ctx, deviceID)
	if err != nil {
		if err == repository.ErrDeviceNotFound {
			return nil, ErrDeviceNotFound
		}
		return nil, err
	}

	contacts, err := s.grpc.GetClientContact(ctx, &proto.ClientdataRequest{
		SenderJid: device.SenderJID,
	})
	if err != nil {
		return nil, err
	}

	var contactList []model.Contact
	for _, contact := range contacts.GetContacts() {
		phone := util.ExtractPhoneNumber(contact.Jid)
		contactList = append(contactList, model.Contact{
			Name:  contact.Short,
			Phone: phone,
			Short: contact.Name,
		})
	}

	return &model.Contacts{
		DeviceId:      device.DeviceID,
		DeviceAlias:   device.DeviceAlias,
		DeviceName:    device.DeviceName,
		ConnectStatus: device.ConnectStatus,
		Contacts:      contactList,
	}, nil
}

func (s *service) GetGroups(ctx context.Context, deviceID string) (*model.Groups, error) {

	if deviceID == "" {
		return nil, ErrDeviceIDEmpty
	}

	// Cek device terlebih dahulu
	device, err := s.repo.FindDeviceByID(ctx, deviceID)
	if err != nil {
		if err == repository.ErrDeviceNotFound {
			return nil, ErrDeviceNotFound
		}
		return nil, err
	}

	groups, err := s.grpc.GetClientGroup(ctx, &proto.ClientdataRequest{
		SenderJid: device.SenderJID,
	})
	if err != nil {
		return nil, err
	}

	var groupList []model.Group
	for _, group := range groups.GetGroups() {
		groupList = append(groupList, model.Group{
			Name:  group.Name,
			Phone: util.ExtractPhoneNumber(group.Jid),
			Short: group.Short,
		})
	}

	return &model.Groups{
		DeviceId:    device.DeviceID,
		DeviceAlias: device.DeviceAlias,
		DeviceName:  device.DeviceName,
		Groups:      groupList,
	}, nil
}
