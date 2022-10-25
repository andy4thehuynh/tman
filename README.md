# tman
A lightweight, simple Tmux session manager for the casual Tmux user.

Tman provides dead simple configuration with example configurations to get you started. Write your config in TOML once and tman will manage your sessions seamlessly. 

# Inspiration
Tmuxinator is the bulldozer you need when you're a big time tmux user. Think heavy Docker configurations with lots of setup and teardown. I find it bloated for more product driven development. Thus the creation of tman - the shovel! This shovel is meant to be fast and set up your workspace consistently.

# Recommendations
Alias `tman` executable to `tm` for quicker a more succinct command. Add this to your `.zshrc` file:
```
alias -g tm="tman"
```

# Features
- **start** a new tman project provided an existing TOML file for project
```
$ tman start example
```
