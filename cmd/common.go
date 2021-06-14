package cmd

import (
	"errors"

	"github.com/hyperjiang/translate/client"
	"github.com/hyperjiang/translate/translator"
	"github.com/rs/zerolog/log"
)

type options struct {
	inputFile  string
	outputFile string
	sourceLang string
	targetLang string
	listLang   bool
}

func exit(err error) {
	log.Fatal().Err(err).Msg("")
}

func translate(client client.Client, opts options) error {
	var jsonTranslator = translator.NewJSONTranslator(client)
	if err := jsonTranslator.ParseFile(opts.inputFile); err == nil {
		if err := jsonTranslator.Translate(opts.sourceLang, opts.targetLang); err != nil {
			return err
		}
		if err := jsonTranslator.SaveResult(opts.outputFile); err != nil {
			return err
		}
		return nil
	}

	var yamlTranslator = translator.NewYAMLTranslator(client)
	if err := yamlTranslator.ParseFile(opts.inputFile); err == nil {
		if err := yamlTranslator.Translate(opts.sourceLang, opts.targetLang); err != nil {
			return err
		}
		if err := yamlTranslator.SaveResult(opts.outputFile); err != nil {
			return err
		}
		return nil
	}

	return errors.New("unknown file format")
}