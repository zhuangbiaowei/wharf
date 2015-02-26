package routers

import (
	"github.com/astaxie/beego"

	"github.com/dockercn/wharf/controllers"
)

func init() {
	//Web Interface
	beego.Router("/", &controllers.WebController{}, "get:GetIndex")
	beego.Router("/auth", &controllers.WebController{}, "get:GetAuth")
	beego.Router("/setting", &controllers.WebController{}, "get:GetSetting")
	beego.Router("/dashboard", &controllers.WebController{}, "get:GetDashboard")

	beego.Router("/admin/auth", &controllers.WebController{}, "get:GetAdminAuth")
	beego.Router("/admin", &controllers.WebController{}, "get:GetAdmin")

	beego.Router("/u/:namespace/:repository", &controllers.WebController{}, "get:GetRepository")

	//Web API
	web := beego.NewNamespace("/w1",
		beego.NSRouter("/signin", &controllers.UserWebAPIV1Controller{}, "post:Signin"),
		beego.NSRouter("/signup", &controllers.UserWebAPIV1Controller{}, "post:Signup"),
		beego.NSRouter("/profile", &controllers.UserWebAPIV1Controller{}, "get:GetProfile"),
		beego.NSRouter("/namespaces", &controllers.UserWebAPIV1Controller{}, "get:GetNamespaces"),

		//repository routers
		beego.NSRouter("/repository", &controllers.RepoWebAPIV1Controller{}, "post:PostRepository"),

		//team routers
		beego.NSRouter("/users/:username", &controllers.UserWebAPIV1Controller{}, "get:GetUser"),
		beego.NSRouter("/team", &controllers.TeamWebV1Controller{}, "post:PostTeam"),

		//organization routers
		beego.NSRouter("/organizations", &controllers.OrganizationWebV1Controller{}, "get:GetOrganizations"),
		beego.NSRouter("/organization", &controllers.OrganizationWebV1Controller{}, "post:PostOrganization"),
		beego.NSRouter("/organization", &controllers.OrganizationWebV1Controller{}, "put:PutOrganization"),
		beego.NSRouter("/organizations/:org", &controllers.OrganizationWebV1Controller{}, "get:GetOrganizationDetail"),
	)

	//Docker Registry API V1 remain
	beego.Router("/_ping", &controllers.PingAPIV1Controller{}, "get:GetPing")

	//Docker Registry API V1
	apiv1 := beego.NewNamespace("/v1",
		beego.NSRouter("/_ping", &controllers.PingAPIV1Controller{}, "get:GetPing"),
		beego.NSRouter("/users", &controllers.UserAPIV1Controller{}, "get:GetUsers"),
		beego.NSRouter("/users", &controllers.UserAPIV1Controller{}, "post:PostUsers"),

		beego.NSNamespace("/repositories",
			beego.NSRouter("/:namespace/:repo_name/tags/:tag", &controllers.RepoAPIV1Controller{}, "put:PutTag"),
			beego.NSRouter("/:namespace/:repo_name/images", &controllers.RepoAPIV1Controller{}, "put:PutRepositoryImages"),
			beego.NSRouter("/:namespace/:repo_name/images", &controllers.RepoAPIV1Controller{}, "get:GetRepositoryImages"),
			beego.NSRouter("/:namespace/:repo_name/tags", &controllers.RepoAPIV1Controller{}, "get:GetRepositoryTags"),
			beego.NSRouter("/:namespace/:repo_name", &controllers.RepoAPIV1Controller{}, "put:PutRepository"),
		),

		beego.NSNamespace("/images",
			beego.NSRouter("/:image_id/ancestry", &controllers.ImageAPIV1Controller{}, "get:GetImageAncestry"),
			beego.NSRouter("/:image_id/json", &controllers.ImageAPIV1Controller{}, "get:GetImageJSON"),
			beego.NSRouter("/:image_id/layer", &controllers.ImageAPIV1Controller{}, "get:GetImageLayer"),
			beego.NSRouter("/:image_id/json", &controllers.ImageAPIV1Controller{}, "put:PutImageJSON"),
			beego.NSRouter("/:image_id/layer", &controllers.ImageAPIV1Controller{}, "put:PutImageLayer"),
			beego.NSRouter("/:image_id/checksum", &controllers.ImageAPIV1Controller{}, "put:PutChecksum"),
		),
	)

	//Dockerfile Build API V1
	buildv1 := beego.NewNamespace("/b1",
		beego.NSRouter("/build", &controllers.BuilderAPIV1Controller{}, "post:PostBuild"),
		beego.NSRouter("/status", &controllers.BuilderAPIV1Controller{}, "get:GetStatus"),
	)

	beego.AddNamespace(web)
	beego.AddNamespace(apiv1)
	beego.AddNamespace(buildv1)
}