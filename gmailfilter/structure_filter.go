package gmailfilter

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"google.golang.org/api/gmail/v1"
)

func setFilterValuesToState(d *schema.ResourceData, filter *gmail.Filter) error {
	if err := d.Set("action", []interface{}{flattenFilterAction(filter.Action)}); err != nil {
		return fmt.Errorf("error setting action: %s", err)
	}
	if err := d.Set("criteria", []interface{}{flattenFilterCriteria(filter.Criteria)}); err != nil {
		return fmt.Errorf("error setting criteria: %s", err)
	}
	return nil
}

func expandFilterAction(d *schema.ResourceData) *gmail.FilterAction {
	raw, ok := d.Get("action").([]interface{})
	if !ok || len(raw) == 0 {
		return nil
	}

	in, ok := raw[0].(map[string]interface{})
	if !ok {
		return nil
	}

	return &gmail.FilterAction{
		AddLabelIds:    expandStringSlice(in["add_label_ids"]),
		Forward:        in["forward"].(string),
		RemoveLabelIds: expandStringSlice(in["remove_label_ids"]),
	}
}

func flattenFilterAction(action *gmail.FilterAction) map[string]interface{} {
	return map[string]interface{}{
		"add_label_ids":    action.AddLabelIds,
		"forward":          action.Forward,
		"remove_label_ids": action.RemoveLabelIds,
	}
}

func flattenFilterCriteria(criteria *gmail.FilterCriteria) map[string]interface{} {
	return map[string]interface{}{
		"exclude_chats":   criteria.ExcludeChats,
		"from":            criteria.From,
		"has_attachment":  criteria.HasAttachment,
		"negated_query":   criteria.NegatedQuery,
		"query":           criteria.Query,
		"size":            criteria.Size,
		"size_comparison": criteria.SizeComparison,
		"subject":         criteria.Subject,
		"to":              criteria.To,
	}
}

func expandFilterCriteria(d *schema.ResourceData) *gmail.FilterCriteria {
	raw, ok := d.Get("criteria").([]interface{})
	if !ok || len(raw) == 0 {
		return nil
	}

	in, ok := raw[0].(map[string]interface{})
	if !ok {
		return nil
	}

	return &gmail.FilterCriteria{
		ExcludeChats:   in["exclude_chats"].(bool),
		From:           in["from"].(string),
		HasAttachment:  in["has_attachment"].(bool),
		NegatedQuery:   in["negated_query"].(string),
		Query:          in["query"].(string),
		Size:           int64(in["size"].(int)),
		SizeComparison: in["size_comparison"].(string),
		Subject:        in["subject"].(string),
		To:             in["to"].(string),
	}
}
