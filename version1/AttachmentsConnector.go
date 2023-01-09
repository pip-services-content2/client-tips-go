package version1

import (
	"context"

	aclients "github.com/pip-services-content2/client-attachments-go/version1"
)

type AttachmentsConnector struct {
	attachmentsClient aclients.IAttachmentsClientV1
}

func NewAttachmentsConnector(client aclients.IAttachmentsClientV1) *AttachmentsConnector {
	return &AttachmentsConnector{
		attachmentsClient: client,
	}
}

func (c *AttachmentsConnector) extractAttachmentIds(tip *TipV1) []string {
	ids := make([]string, 0)

	for _, pic := range tip.Pics {
		if pic.Id != "" {
			ids = append(ids, pic.Id)
		}
	}

	for _, doc := range tip.Docs {
		if doc.Id != "" {
			ids = append(ids, doc.Id)
		}
	}

	return ids
}

func (c *AttachmentsConnector) AddAttachments(ctx context.Context, correlationId string, tip *TipV1) error {
	if c.attachmentsClient == nil || tip == nil {
		return nil
	}

	ids := c.extractAttachmentIds(tip)
	reference := aclients.NewReferenceV1(tip.Id, "tip", "")
	_, err := c.attachmentsClient.AddAttachments(ctx, correlationId, reference, ids)
	return err
}

func (c *AttachmentsConnector) UpdateAttachments(ctx context.Context, correlationId string, oldTip *TipV1, newTip *TipV1) error {
	if c.attachmentsClient == nil || newTip == nil || oldTip == nil {
		return nil
	}

	oldIds := c.extractAttachmentIds(oldTip)
	newIds := c.extractAttachmentIds(newTip)
	reference := aclients.NewReferenceV1(newTip.Id, "tip", "")
	_, err := c.attachmentsClient.UpdateAttachments(ctx, correlationId, reference, oldIds, newIds)

	return err
}

func (c *AttachmentsConnector) RemoveAttachments(ctx context.Context, correlationId string, tip *TipV1) error {
	if c.attachmentsClient == nil || tip == nil {
		return nil
	}

	ids := c.extractAttachmentIds(tip)
	reference := aclients.NewReferenceV1(tip.Id, "tip", "")
	_, err := c.attachmentsClient.RemoveAttachments(ctx, correlationId, reference, ids)

	return err
}
