package templates

import (
	"fmt"

	"sigs.k8s.io/kubebuilder/v3/pkg/model/file"
)

var _ file.Template = &OperatorFile{}

type OperatorFile struct {
	file.TemplateMixin

	// Package is the source files package
	Package string

	// Name of the operator used for the main file.
	OperatorName string
}

func (f *OperatorFile) SetTemplateDefaults() error {
	if f.OperatorName == "" {
		return fmt.Errorf("invalid operator name")
	}

	if f.Path == "" {
		f.Path = f.OperatorName + "Operator.java"
	}

	f.TemplateBody = operatorTemplate

	return nil
}

// TODO: pass in the name of the operator i.e. replace Memcached
const operatorTemplate = `
package {{ .Package }};

import io.javaoperatorsdk.operator.Operator;
import io.quarkus.runtime.Quarkus;
import io.quarkus.runtime.QuarkusApplication;
import io.quarkus.runtime.annotations.QuarkusMain;
import javax.inject.Inject;

@QuarkusMain
public class {{ .OperatorName }}Operator implements QuarkusApplication {

  @Inject Operator operator;

  public static void main(String... args) {
    Quarkus.run({{ .OperatorName }}Operator.class, args);
  }

  @Override
  public int run(String... args) throws Exception {
    operator.start();

    Quarkus.waitForExit();
    return 0;
  }
}
`
