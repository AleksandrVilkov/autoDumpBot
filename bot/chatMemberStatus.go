package bot

type ChatMemberStatus string

const (
	MEMBER  ChatMemberStatus = "member"        //пользователь является подписчиком;
	LEFT                     = "left"          //пользователь не подписан;
	KICKED                   = "kicked"        // пользователь заблокирован;
	ADMIN                    = "administrator" //админ
	CREATOR                  = "creator"       //создатель
)
