# Docktopia

```
     _            _    _              _
  __| | ___   ___| | _| |_ ___  _ __ (_) __ _
 / _` |/ _ \ / __| |/ / __/ _ \| '_ \| |/ _` |
| (_| | (_) | (__|   <| || (_) | |_) | | (_| |
 \__,_|\___/ \___|_|\_\\__\___/| .__/|_|\__,_|
                               |_|
```

`docktopia` is a user-friendly command-line tool designed to simplify the management of Docker-based projects, including creating and managing required files such as `docker-compose.yml`, `Dockerfile`, and other necessary configuration files. With `docktopia`, you can easily define, configure, and operate multiple Docker projects with ease.

At the core of `docktopia` lies the concept of crates, which serve as self-contained units for managing and organizing collections of interconnected Docker projects. Each crate consists of several components called targets, representing the individual Docker projects within the crate. By structuring projects into crates and targets, `docktopia` enables efficient organization, maintenance, and operation of even the most intricate multi-project environments.

A significant advantage of `docktopia` is its extensibility through hooks – customizable scripts or commands that execute during specific phases of various processes managed by `docktopia`. These hooks empower users to add their own logic and automate tasks according to their unique needs. They offer flexibility and enable seamless integration with external systems, allowing for advanced configurations and workflows.

## Key Features

- **Crates and Targets:** `docktopia` provides a flexible and organized structure for managing Docker-based projects. Crates represent a collection of targets that can be used to control the operations for all of its targets via hooks. It can also hold configuration and other files used across most or all of its targets. The targets also provide their own set of hooks and are the structural elements actually used to store and manage projects.

- **Hooks Support:** Both crates and targets support hooks, allowing you to run custom scripts or commands. These hooks provide a modular way to extend and customize the operations according to your specific needs. With hooks, you can seamlessly integrate additional functionality, perform complex transformations, or interact with external systems.

- **Templates Support:** Creating new crates and targets is made easier with template support. Templates provide a convenient way to create consistent configurations by predefining common settings and hooks.

- **Utilities for Hooks:** `docktopia` includes a set of utilities that can be used with hooks to simplify common tasks. These utilities allow you to display messages, show section titles, print colored output, ask for user confirmation, and more. They enhance the functionality of hooks and provide a convenient way to interact with the user during operations.

- **Command-Line Interface and User-Friendly Structure:** With its command-line interface, `docktopia` offers a straightforward and efficient way to manage your operations. The program is designed with a user-friendly structure, providing intuitive commands and comprehensive documentation.

## Installation

To build the program, make sure that Go is installed on your system. Clone the repository or download an archive for a specific version and run the following command in the terminal:

```bash
make build
```

This will create a binary file called `docktopia`; autocompletion files for `bash`, `zsh`, and `fish` in the directory `./autocompletions`; and markdown documentation files in the directory `./docs`. To install the binary and program files (autocompletion files and documentation):

```bash
sudo make install
```

And to uninstall:

```bash
sudo make uninstall
```

To remove files and directories created by this operation:

```bash
make clean
```

### Custom Destination Directory

By default, all files will be installed to their respectives subdirectories under the `/usr` directory in your root filesystem. However, if you want to set a custom destination directory for the installation, you can use the `DESTDIR` variable when running the `make install` command (also valid for `make uninstall`). For example, here's how you could build this program for Termux:

```bash
make install DESTDIR=/data/data/com.termux/files/usr
```

### PKGBUILD

To simplify the building and installation process on Arch Linux, a PKGBUILD is available in this repository.

To build the package, run the following command in the terminal:
```bash
makepkg -sf
```

If you want to build a package specifically for Termux when using `pacman` as the package manager instead of the default one (`apt`), use the following command:
```bash
TERMUX_BUILD= makepkg -sf
```

To ensure that the PKGBUILD correctly identifies Termux as the build target and sets the installation destination directory correctly, the `TERMUX_BUILD` environment variable should be set. It doesn't require any specific value; simply setting the variable is sufficient for the PKGBUILD to recognize Termux as the build target and handle the installation directory accordingly.

> For more information on how to switch package managers on Termux, refer to the [Termux Wiki](https://wiki.termux.com/wiki/Switching_package_manager).

### Initializing User Data Directory

Before you begin using `docktopia`, you should create the user data directory by executing the following command:

```bash
docktopia init
```

By default, this command will generate the necessary directory structure under `~/docktopia`. However, if you prefer a different location for your data directory, you can specify a custom path using the `-D` flag as shown below:

```bash
docktopia init -D <custom_data_dir>
```

## Usage

The general usage of Docktopia is as follows:

```
docktopia [command]
```

To get started, run the following command:

```
docktopia
```

This will display information about the program and provide instructions on how to proceed. To learn more about the available commands and options, you can use the `--help/-h` flag:

```
docktopia --help
```

This command will provide detailed information about the available commands, their usage, and the available options for each command. It is a useful reference when you need more information on how to use a specific command or what options are available for customization.

Feel free to explore the available commands and options using the `--help` flag to get a better understanding of the functionality provided by Docktopia.

### Documentation

To generate the program's documentation in markdown format, you can use the following command:

```shell
docktopia docs generate
```

The documentation will be generated in the `./docs` directory by default. You can specify a custom location by using:

```shell
docktopia docs generate -o <output_dir>
```

## Documentation

### Available Subcommands

- docktopia: The main command for the program.
  - version: Show the program's version.
  - init: Create user data directory
  - docs: Program documentation.
    - generate: Generate program documentation (markdown files).
  - utils: Utilities for hooks execution.
    - msg: Print a message in a specific color given a HEX code.
    - attention: Display attention message.
    - error: Display error message.
    - success: Display success message.
    - section: Display section title.
    - hr: Display a horizontal line.
    - confirm: Ask for confirmation.
    - sshagent-start: Starts an ssh-agent process and adds a private key.
    - sshagent-stop: Stops the ssh-agent process.
    - sshagent-getpid: Get the process ID of the ssh-agent.
    - sshagent-getsock: Get the socket path of the ssh-agent.
  - crates: Manage crates.
    - edit: Edit crates.
    - view: View crates.
    - enable: Enable crates.
    - disable: Disable crates.
    - create: Create crates.
    - rm: Remove crates.
    - ls: List crates.
    - hooks: Manage crate hooks.
      - run: Run crate hook(s).
      - ls: List crate hooks.
  - targets: Manage targets.
    - ls: List targets.
    - edit: Edit targets.
    - view: View targets.
    - enable: Enable targets.
    - disable: Disable targets.
    - create: Create targets.
    - rm: Remove targets.
    - hooks: Manage target hooks.
      - run: Run target hook(s).
      - ls: List target hooks.

### User Data Directory

The default user data directory is located at `~/docktopia`. Its tree structure is as follows:

- `templates`: This directory contains templates used for creating crates and targets. It provides a starting point with pre-configured setups for common scenarios. The templates are organized into subdirectories based on their type, such as `crates` and `targets`.

- `templates/crates`: This subdirectory within the `templates` directory contains templates specifically designed for creating crates. Each template may include a set of pre-defined hook scripts and configuration files to streamline the crate creation process.

- `templates/targets`: This subdirectory within the `templates` directory contains templates specifically designed for creating targets. Similar to the crate templates, each target template may include hook scripts and configuration files tailored for specific needs.

- `crates`: This directory holds the configurations and settings for all created crates. Each crate has its own subdirectory within the `crates` directory. The subdirectories are named after the respective crate and contain the associated configuration files, hooks, and any other necessary files.

- `crates/<crate>/targets`: Within each crate's subdirectory, there is a `targets` directory. This directory holds the configurations and hooks for all the targets associated with that particular crate. Each target has its own subdirectory within the `targets` directory, containing the target-specific configuration files, hooks, and any other necessary files.

#### Custom User Data Directory

It is possible to set a custom user data directory using the `-D` flag. To do so, you can run the `docktopia` program followed by the flag and the desired directory path. The command would look like this:

```
docktopia -D <custom_data_dir>
```

This flag can be used with any subcommand of the `docktopia` program. Whether you are creating crates, managing hooks, or performing other operations, you have the flexibility to specify a custom data directory that best suits your needs.

### Crates

`Crates` are the primary structural element in this program. A crate represents a collection of projects (which are called `targets`).

Think of a crate as a container that holds the necessary information to manage projects. Within a crate, you can define one or more targets, each corresponding to a specific project (e.g., a software package) you want to manage.

#### Creating a Crate

To create a new crate, you can utilize the `crates create` command. This command will guide you through the process by prompting for a crate name and allowing you to choose a template. Templates provide pre-configured setups for common scenarios.

If there are no crate templates available or if you prefer to have a minimal setup with just the crate directory and a `targets` subdirectory, you can select the `scratch` template. This template will only create the necessary directories and will not include any additional configuration files or hooks.

After creating the crate, a new directory will be generated specifically for that crate. This directory will contain the relevant configuration files and hooks, based on the selected template.

#### Managing Crates

Once you have created crates, you can perform various operations on them. The following commands are available for managing crates:

- `crates edit`: Edit crates.
- `crates view`: View crates.
- `crates rm`: Remove crates.
- `crates ls`: List crates.
- `crates enable`: Enable crates.
- `crates disable`: Disable crates.
- `crates hooks`: Manage crate hooks.
 - `crates hooks run`: Run crate hook(s).
 - `crates hooks ls`: List crate hooks.


#### Disabled Crates

When executing the `crates hooks run` command or other commands that depend on hooks to function (like `crates edit` and `crates view`), `docktopia` will check if the crate is disabled before running hooks for each crate. If any disabled crates are encountered, `docktopia` will skip them without generating an error. It will proceed to the remaining selected crates, if any are available.

#### Hooks

Hooks are an essential component of crates, allowing you to define custom actions or scripts to execute. The following default hooks can be associated with a crate:

- `post_create`: Runs after a crate is created. Can be used to further configure the crate configuration beyond only creating its directory (which is done automatically by the program).
- `pre_rm`: Executes before removing a crate.
- `ls`: Displays custom information for the crate when running `crates ls`.
- `edit`: Script to open the crate configuration when running `crates edit`.
- `view`: Script to open the crate configuration when running `crates view`.

##### Environment Variables

When running crate hooks, the following environment variables are available for your use:

- **PROGRAM_NAME**: The name of the program (`docktopia`).
- **DEFAULT_SHELL**: The default shell used by the program.
- **DOCKTOPIA_EXEC**: The executable path of `docktopia`.
- **DOCKTOPIA_UTILS**: The executable path of `docktopia utils`, providing access to the utility commands.
- **USER_DATA_DIR**: The user data directory used by `docktopia`.
- **USER_CRATES_DIR**: The directory where user crates are stored.
- **USER_TEMPLATES_DIR**: The directory where user templates are stored.
- **USER_CRATES_TEMPLATES_DIR**: The directory where user crate templates are stored.
- **USER_TARGETS_TEMPLATES_DIR**: The directory where user target templates are stored.
- **CRATE_NAME**: The name of the current crate.
- **CRATE_DIR**: The directory path of the current crate.
- **CRATE_HOOKS_DIR**: The directory path of the hooks within the current crate.
- **CRATE_TARGETS_DIR**: The directory path of the targets within the current crate.
- **CRATE_TEMP_DIR**: The temporary directory path specific to the current crate.

These environment variables provide useful information and paths that can be utilized within your crate hooks to customize the behavior and perform specific actions based on the current context.

##### Custom entry command

A custom entry command for a specific hook can be defined by creating a file called `<hook_name>.entry` in the crate's hooks directory. For example, if you want to run the `post_create` hook with the `fish` shell, you could do so by creating a `post_create.entry` file with the following content:

```bash
/usr/bin/fish
```

##### Managing Hooks

When managing hooks at the crate level, you can use the following subcommands:

- `crates hooks ls`: List all hooks available for a specific crate. This command displays the names of the hooks and their associated entry commands (if any).

- `targets hooks run`: Run one or more hooks for the specified crates. You can also run hooks in a specific sequence by running, for example:

```bash
crates hooks run --crate <crate_name> --hook <first_hook> --hook <second_hook>
```

##### Temporary Directory

When running hooks for crates, a temporary directory is created for each crate. This temporary directory serves as a workspace for performing actions or modifications during the hook execution process. It is called `.tmp` and created within the crate directory itself. This crate-specific temporary directory provides a separate workspace for any temporary files or data specific to the individual crate and can be used for any shared temporary files or data required by the hook scripts across multiple targets within the same crate. It allows the hooks to operate within a context that is isolated to the crate directory.

Once the hook execution is complete, the temporary directory and its contents are typically cleaned up automatically. When running hooks via `crates hooks run`, this behavior can be modified using two options (non-exclusive):

- `nocreatetemp`: By using this option, temporary directories will not be created before running the hook(s). This can be useful if you prefer to handle temporary directory creation manually or if your hook scripts do not require a separate workspace.

- `noremovetemp`: With this option, temporary directories will not be automatically removed after running the hook(s). This allows you to inspect or access the temporary directories and their contents after the hook execution has completed. It can be beneficial for debugging purposes or if you need to access the temporary files generated during the hook execution.

### Targets

`Targets` are the secondary structural element and represent a specific project. Each target is associated with a crate and can have its own set of configurations and hooks.

#### Creating a Target

To create a new target, you can use the `targets create` command. This command will guide you through the process by allowing you to select a parent crate and provide a name for the target. Additionally, you can choose a template to pre-configure the target's setup according to your requirements.

If there are no target templates available or if you prefer a minimal setup for your target, you can select the `scratch` template during the target creation process. This template will create only the target's directory. It will not include any additional configuration files or hooks, providing you with a clean slate to customize according to your needs.

#### Managing Targets

Once you have created targets, you can perform various operations on them. The following commands are available for managing targets:

- `targets ls`: List targets.
- `targets edit`: Edit targets.
- `targets view`: View targets.
- `targets sync`: Sync targets.
- `targets enable`: Enable targets.
- `targets disable`: Disable targets.
- `targets rm`: Remove targets.
- `targets hooks`: Manage target hooks.
 - `targets hooks run`: Run target hook(s).
 - `targets hooks ls`: List target hooks.

#### Disabled Targets

When executing the `targets hooks run` command or other commands that depend on hooks to function (like `targets edit` and `targets view`), `docktopia` will check if the target is disabled before initiating the synchronization process for each target. It will also verify if the crate is disabled before iterating through the targets. If any disabled targets are encountered, `docktopia` will skip them without generating an error. It will proceed to sync the remaining selected targets, if any are available (only if the crate is enabled). In essence, `docktopia` gracefully handles disabled targets during the synchronization process, allowing for the successful synchronization of the remaining enabled targets.

#### Hooks

Similar to crates, targets also support hooks that allow you to define custom actions or scripts to execute at specific stages of the synchronization process. The default hooks available for targets include:

- `post_create`: Runs after a target is created. Can be used to further configure the target configuration beyond only creating its directory (which is done automatically by the program).
- `pre_rm`: Executes before removing a target.
- `ls`: Displays custom information for the target when running `targets ls`.
- `edit`: Script to open the target configuration when running `targets edit`.
- `view`: Script to open the target configuration when running `targets view`.

##### Environment Variables

When running target hooks, the following environment variables are available for your use:

- **PROGRAM_NAME**: The name of the program (`docktopia`).
- **DEFAULT_SHELL**: The default shell used by the program.
- **DOCKTOPIA_EXEC**: The executable path of `docktopia`.
- **DOCKTOPIA_UTILS**: The executable path of `docktopia utils`, providing access to the utility commands.
- **USER_DATA_DIR**: The user data directory used by `docktopia`.
- **USER_CRATES_DIR**: The directory where user crates are stored.
- **USER_TEMPLATES_DIR**: The directory where user templates are stored.
- **USER_CRATES_TEMPLATES_DIR**: The directory where user crate templates are stored.
- **USER_TARGETS_TEMPLATES_DIR**: The directory where user target templates are stored.
- **CRATE_NAME**: The name of the current crate.
- **CRATE_DIR**: The directory path of the current crate.
- **CRATE_HOOKS_DIR**: The directory path of the hooks within the current crate.
- **CRATE_TARGETS_DIR**: The directory path of the targets within the current crate.
- **CRATE_TEMP_DIR**: The temporary directory path specific to the current crate.
- **TARGET_NAME**: The name of the current target.
- **TARGET_DIR**: The directory path of the current target.
- **TARGET_HOOKS_DIR**: The directory path of the hooks within the current target.
- **TARGET_TEMP_DIR**: The temporary directory path specific to the current target.

These environment variables provide useful information and paths that can be utilized within your target hooks to customize the behavior and perform specific actions based on the current context.

##### Custom entry command

A custom entry command for a specific hook can be defined by creating a file called `<hook_name>.entry` in the target's hooks directory. For example, if you want to run the `post_create` hook with the `fish` shell, you could do so by creating a `post_create.entry` file with the following content:

```bash
/usr/bin/fish
```
##### Managing Hooks

When working with targets, you have the option to use the following subcommands to manage their hooks:

- `targets hooks ls`: This command lists all available hooks for the specified targets. It displays the names of the hooks and their associated entry commands (if any).

- `targets hooks run`: Use this command to run one or more hooks for the specified targets. You can also specify a specific sequence for the hooks by running the command as follows:

```bash
targets hooks run --crate <crate_name> --target <target_name> --hook <first_hook> --hook <second_hook>
```

Additionally, the `targets hooks run` command allows you to specify an array of crate hooks to be executed before and after running the targets' hooks. You can achieve this by using the following two flags:

- `cratepre`: Specifies the hooks to be run before iterating through the targets.
- `cratepost`: Specifies the hooks to be run after iterating through the targets.

Similar to the `--hooks/-k` flag, you can add multiple flags for both options. The hooks will be executed in the sequence they were specified in the command line. For example, you could run the following command:

```bash
targets hooks run --crate <crate_name> --target <target_name> --cratepre pre_build --cratepost post_build --hook build --hook install --hook clean
```

##### Temporary Directory

When running hooks for targets, two temporary directories are created. These temporary directories serve as a workspace for performing actions or modifications during the hook execution process.

- The first one is the `.tmp` directory within the crate directory, similar to the crate hooks. This temporary directory is used for any shared temporary files or data required by the hook scripts across multiple targets within the same crate.

- Additionally, a second temporary directory called `.tmp` is created within the target directory itself. This target-specific temporary directory provides a separate workspace for any temporary files or data specific to the individual target. It allows the hooks to operate within a context that is isolated to the target directory.

Once the hook execution is complete, the temporary directories and their contents are typically cleaned up automatically. When running hooks via `targets hooks run`, this behavior can be modified using two options (non-exclusive):

- `nocreatetemp`: By using this option, temporary directories will not be created before running the hook(s). This can be useful if you prefer to handle temporary directory creation manually or if your hook scripts do not require a separate workspace.

- `noremovetemp`: With this option, temporary directories will not be automatically removed after running the hook(s). This allows you to inspect or access the temporary directories and their contents after the hook execution has completed. It can be beneficial for debugging purposes or if you need to access the temporary files generated during the hook execution.

### Utilities

`docktopia` provides a set of utilities designed to be used within the hooks of crates and targets, allowing you to perform additional actions or execute custom logic during operations, though they can be used wherever and whenever you want. The main difference is that when running crate and target hooks, an environment variable called `$DOCKTOPIA_UTILS` is automatically created, pointing to `docktopia utils`.

#### Available Utilities

- **msg**: Print a message in a specific color given a HEX code.
  ```
  docktopia utils msg '#FF5733' 'This is a colored message'
  ```

- **attention**: Display an attention message.
  ```
  docktopia utils attention 'This is an attention message'
  ```

- **error**: Display an error message.
  ```
  docktopia utils error 'This is an error message'
  ```

- **success**: Display a success message.
  ```
  docktopia utils success 'This is a success message'
  ```

- **section**: Display a section title.
  ```
  docktopia utils section 'This is a section title'
  ```

- **hr**: Display a horizontal line.
  ```
  docktopia utils hr '-' 0.45
  ```

- **confirm**: Ask for confirmation from the user.
  ```
  docktopia utils confirm 'Are you sure you want to continue?'
  ```

- **sshagent-start**: Start an ssh-agent process and add a private key.
  ```
  docktopia utils sshagent-start '/path/to/my/key' $TARGET_TEMP_DIR
  ```

- **sshagent-stop**: Stop the ssh-agent process.
  ```
  docktopia utils sshagent-stop $TARGET_TEMP_DIR
  ```

- **sshagent-getpid**: Get the process ID of the ssh-agent.
  ```
  docktopia utils sshagent-getpid $TARGET_TEMP_DIR
  ```

- **sshagent-getsock**: Get the socket path of the ssh-agent.
  ```
  docktopia utils sshagent-getsock $TARGET_TEMP_DIR
  ```

#### Using Utilities in Hooks

To use any of the utilities within a crate or target hook, you can access them using the `$DOCKTOPIA_UTILS` environment variable, which points to the command `docktopia utils`. For example, to display an attention message within a crate hook:

```bash
$DOCKTOPIA_UTILS attention 'This is an attention message'
```

This will print the specified attention message to the console, attracting the user's attention during the operation.

### Templates

Templates play a crucial role in customizing the creation of crates and targets. `docktopia` provides a template-based approach to create crates and targets, allowing you to quickly set up and configure your project structure. When creating a crate or target, `docktopia` automatically generates the corresponding crate or target directory and copies the selected template structure to it.
To simplify template usage, you can find a collection of example crate and target templates that I personally use in the `./templates` directory within the source code. These templates serve as starting points and can be customized to suit your specific project requirements. Additionally, you have the flexibility to create your own templates and store them in the following directories within the user data directory:

- `templates/crates`: This directory is dedicated to crate templates.
- `templates/targets`: This directory is dedicated to target templates.

By placing your custom templates in these directories, they become readily available for selection during the crate and target creation process. You can leverage these templates to expedite the setup of your projects and tailor them to your specific needs.

## License

Docktopia is licensed under the GPL-3.0 license.
