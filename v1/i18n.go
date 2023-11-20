package cli

type (
	I18nTag string
)

var (
	I18n = map[I18nTag]string{}

	Tag_Default I18nTag = "default"

	Tag_OutputCmdHelp    I18nTag = "output_cmd_help"
	Tag_OutputSubCmdHelp I18nTag = "output_sub_cmd_help"

	Tag_OutputVersion I18nTag = "output_version"

	Tag_HelpDescription    I18nTag = "help_description"
	Tag_VersionDescription I18nTag = "version_description"
)

func UseEN() {
	I18n[Tag_OutputCmdHelp] = template_en_OutputCmdHelp
	I18n[Tag_OutputSubCmdHelp] = template_en_OutputSubCmdHelp
	I18n[Tag_OutputVersion] = template_en_OutputVersion

	I18n[Tag_HelpDescription] = "show help"
	I18n[Tag_VersionDescription] = "show curr version"

	I18n[Tag_Default] = "default"
}
