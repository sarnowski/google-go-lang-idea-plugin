package ro.redeul.google.go.runner;

import com.intellij.execution.configurations.ConfigurationFactory;
import com.intellij.execution.configurations.ConfigurationType;
import com.intellij.execution.configurations.RunConfiguration;
import com.intellij.openapi.extensions.Extensions;
import com.intellij.openapi.project.Project;
import com.intellij.util.containers.ContainerUtil;
import org.jetbrains.annotations.NonNls;
import org.jetbrains.annotations.NotNull;
import ro.redeul.google.go.GoIcons;

import javax.swing.*;

/**
 * Created by IntelliJ IDEA.
 * User: mtoader
 * Date: Aug 19, 2010
 * Time: 2:49:26 PM
 * To change this template use File | Settings | File Templates.
 */
public class GoRunConfigurationType implements ConfigurationType {

    private final GoFactory myConfigurationFactory;

    public GoRunConfigurationType() {
        myConfigurationFactory = new GoFactory(this);
    }

    public String getDisplayName() {
        return "Go Application";
    }

    public String getConfigurationTypeDescription() {
        return "Go Application";
    }

    public Icon getIcon() {
        return GoIcons.GO_ICON_16x16;
    }

    @NonNls
    @NotNull
    public String getId() {
        return "GoApplicationRunConfiguration";
    }

    public ConfigurationFactory[] getConfigurationFactories() {
        return new ConfigurationFactory[]{myConfigurationFactory};
    }

    public static GoRunConfigurationType getInstance() {
        return ContainerUtil.findInstance(Extensions.getExtensions(CONFIGURATION_TYPE_EP), GoRunConfigurationType.class);
    }

    public static class GoFactory extends ConfigurationFactory {

        public GoFactory(ConfigurationType type) {
            super(type);
        }

        public RunConfiguration createTemplateConfiguration(Project project) {
            return new GoApplicationConfiguration("Go application", project, getInstance());
        }        
    }
}
