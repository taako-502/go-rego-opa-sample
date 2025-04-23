package demo

# デフォルトでは拒否
default allow = false

# 管理者は無条件で許可
allow = true if input.user.role == "admin"


# developers グループのメンバーはどのアクションでも許可
allow = true if input.user.groups[_] == "developers"

# 18歳以上かつ action が "read" の場合に許可
allow = true if {
  input.user.age >= 18
  input.action == "read"
}
