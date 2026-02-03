"# LearnGolang" 


Debug

```
   {
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch main",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}"
        },
        {
            "name": "Launch current file",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${file}"
        },
        {
            "name": "Launch current package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}"
        },
        {
            "name": "Attach to Process",
            "type": "go",
            "request": "attach",
            "mode": "local",
            "processId": 0
        }
    ]
}

```