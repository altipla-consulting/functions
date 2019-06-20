
workflow "Test on push" {
  on = "push"
  resolves = ["Test"]
}

action "Test" {
  uses = "actions-contrib/go@master"
  args = "test -race ./..."
}
