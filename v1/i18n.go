package cli

type (
	I18nTag string
)

var (
	I18n = map[I18nTag]string{}

	tag_OutputCmdHelp    I18nTag = "output_cmd_help"
	tag_OutputSubCmdHelp I18nTag = "output_sub_cmd_help"

	tag_HelpDescription    I18nTag = "help_description"
	tag_VersionDescription I18nTag = "version_description"
)

func UseEN() {
	I18n[tag_OutputCmdHelp] = template_en_OutputCmdHelp
	I18n[tag_OutputSubCmdHelp] = template_en_OutputSubCmdHelp

	I18n[tag_HelpDescription] = "show help"
	I18n[tag_VersionDescription] = "show curr version"
}
