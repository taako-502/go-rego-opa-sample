package demo

default allow = false

# Allow admins unconditionally
allow {
    input.user.role == "admin"
}

# Allow users in 'developers' group for any action
allow {
    input.user.groups[_] == "developers"
}

# Allow users age >= 18 to perform 'read' action
allow {
    input.user.age >= 18
    input.action == "read"
}