.goat/
│
├── HEAD
├── config
├── description
├── hooks/
│   └── [sample hook scripts]
├── info/
│   └── exclude
├── objects/
│   └── [object data and directories]
├── refs/
│   ├── heads/
│   │   └── main
│   └── tags/
├── index
├── packed-refs
├── logs/
│   └── refs/
│       └── heads/
│           └── main
├── COMMIT_EDITMSG
└── ORIG_HEAD

.goat/
├── HEAD                   # Points to 'ref: refs/heads/main'
├── refs/
│   ├── heads/
│   │   └── main           # Contains the latest commit hash on the 'main' branch
│   └── tags/              # Contains tag references (if any)
├── logs/
│   └── refs/
│       └── heads/
│           └── main       # Contains the commit history of the 'main' branch
