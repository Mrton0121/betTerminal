![](./assets/banner.gif)

# betTerminal
A linux CLI tool made in go to replace bashrc aliases and to easily write macros in Linux.

## Try out the latest install with the following command:
```sh
bash -c "$(curl -fsSL https://raw.githubusercontent.com/Mrton0121/betTerminal/refs/heads/installation-script/install.sh)"
```

## Usage

The whole tool is based around a single yaml file

This yaml file describes every command that you want to use

These are built-in linux commands with placeholders for arguments, and they will be evaluated by linux itself

You can check out the **example-conf.yaml** for pointers.

The yaml structure is the following with an example: 
```yaml
commands:
  - name: cd # name of the command that you will be using (e.g.: bt cd yourfolder)
    alias: # a list of alternative names for the command
      - changeDir 
    argCount: 1 # number of the arguments that the command will take, if provided more or less then the command will fail and do nothing
    exec: # list of linux commands that will be executed sequentially 
      - echo $(pwd) > ~/.betterminal/lastdir
      - cd $1 # the argument is placed in the place of the $1. The n-th argument that you provide will be placed in $n
    helpText: "Arguments: \n\t1: the path of the directory you want to go" # helping text if there are less or more arguments
```

You can even write commands that use other **bt** commands or even provide parameters for **bt** itself

In this example, it takes the builtin historys n-1th item and uses that as the new command (the last would be the _redo_ command)
```yaml
  - name: redo
    argCount: 0
    exec:
      - bt $(tail -n 2 $BETTERMINAL_CONFIG/history | head -n 1)
```

## Built-in features

After executing a command it will get written into the history file in your **BETTERMINAL_CONFIG** path _($HOME/.betterminal on default)_

You have a **bt init** command that can download any raw yaml file from the web making your config sharing easy

Try it out with the following:
```sh
bt init https://raw.githubusercontent.com/Mrton0121/betTerminal/refs/heads/main/example-conf.yaml
```
This will install the _example-conf.yaml_ into your **bt instance**
