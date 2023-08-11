package delegates

import (
	"reflect"
	"testing"
)

func TestMockDelegate_New(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockRole                 *Role
		MockAddUserAuthResp      *ResponseID
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockAddLikeRes           *ResponseID
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockActivateCommentRes   *Response
		MockDeActivateCommentRes *Response
		MockUnactivatedUserList  *[]User
		MockBannedUserList       *[]User
		MockEnableUser           *Response
		MockReactivateBannedUser *Response
		MockDisableUser          *Response
		MockDisableUserForCause  *Response
		MockConfig               *Config
		MockUpdateConfigRes      *Response
	}
	tests := []struct {
		name   string
		fields fields
		want   Delegate
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MockDelegate{
				MockAddUserResp:          tt.fields.MockAddUserResp,
				MockRole:                 tt.fields.MockRole,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockActivateCommentRes:   tt.fields.MockActivateCommentRes,
				MockDeActivateCommentRes: tt.fields.MockDeActivateCommentRes,
				MockUnactivatedUserList:  tt.fields.MockUnactivatedUserList,
				MockBannedUserList:       tt.fields.MockBannedUserList,
				MockEnableUser:           tt.fields.MockEnableUser,
				MockReactivateBannedUser: tt.fields.MockReactivateBannedUser,
				MockDisableUser:          tt.fields.MockDisableUser,
				MockDisableUserForCause:  tt.fields.MockDisableUserForCause,
				MockConfig:               tt.fields.MockConfig,
				MockUpdateConfigRes:      tt.fields.MockUpdateConfigRes,
			}
			if got := d.New(); got == nil {
				t.Errorf("MockDelegate.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_AddUser(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockRole                 *Role
		MockAddUserAuthResp      *ResponseID
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockAddLikeRes           *ResponseID
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockActivateCommentRes   *Response
		MockDeActivateCommentRes *Response
		MockUnactivatedUserList  *[]User
		MockBannedUserList       *[]User
		MockEnableUser           *Response
		MockReactivateBannedUser *Response
		MockDisableUser          *Response
		MockDisableUserForCause  *Response
		MockConfig               *Config
		MockUpdateConfigRes      *Response
	}
	type args struct {
		u *User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ResponseID
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockAddUserResp: &ResponseID{
					ID:      4,
					Success: true,
				},
			},
			args: args{
				u: &User{},
			},
			want: &ResponseID{
				ID:      4,
				Success: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MockDelegate{
				MockAddUserResp:          tt.fields.MockAddUserResp,
				MockRole:                 tt.fields.MockRole,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockActivateCommentRes:   tt.fields.MockActivateCommentRes,
				MockDeActivateCommentRes: tt.fields.MockDeActivateCommentRes,
				MockUnactivatedUserList:  tt.fields.MockUnactivatedUserList,
				MockBannedUserList:       tt.fields.MockBannedUserList,
				MockEnableUser:           tt.fields.MockEnableUser,
				MockReactivateBannedUser: tt.fields.MockReactivateBannedUser,
				MockDisableUser:          tt.fields.MockDisableUser,
				MockDisableUserForCause:  tt.fields.MockDisableUserForCause,
				MockConfig:               tt.fields.MockConfig,
				MockUpdateConfigRes:      tt.fields.MockUpdateConfigRes,
			}
			if got := d.AddUser(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.AddUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_UpdateUser(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddUserAuthResp      *ResponseID
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockAddLikeRes           *ResponseID
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockActivateCommentRes   *Response
		MockDeActivateCommentRes *Response
		MockUnactivatedUserList  *[]User
		MockBannedUserList       *[]User
		MockEnableUser           *Response
		MockReactivateBannedUser *Response
		MockDisableUser          *Response
		MockDisableUserForCause  *Response
		MockConfig               *Config
		MockUpdateConfigRes      *Response
	}
	type args struct {
		u *User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockUpdateUserResp: &Response{
					Success: true,
				},
			},
			args: args{
				u: &User{},
			},
			want: &Response{
				Success: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MockDelegate{
				MockAddUserResp:          tt.fields.MockAddUserResp,
				MockUpdateUserResp:       tt.fields.MockUpdateUserResp,
				MockRole:                 tt.fields.MockRole,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockActivateCommentRes:   tt.fields.MockActivateCommentRes,
				MockDeActivateCommentRes: tt.fields.MockDeActivateCommentRes,
				MockUnactivatedUserList:  tt.fields.MockUnactivatedUserList,
				MockBannedUserList:       tt.fields.MockBannedUserList,
				MockEnableUser:           tt.fields.MockEnableUser,
				MockReactivateBannedUser: tt.fields.MockReactivateBannedUser,
				MockDisableUser:          tt.fields.MockDisableUser,
				MockDisableUserForCause:  tt.fields.MockDisableUserForCause,
				MockConfig:               tt.fields.MockConfig,
				MockUpdateConfigRes:      tt.fields.MockUpdateConfigRes,
			}
			if got := d.UpdateUser(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_GetUser(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddUserAuthResp      *ResponseID
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockAddLikeRes           *ResponseID
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockActivateCommentRes   *Response
		MockDeActivateCommentRes *Response
		MockUnactivatedUserList  *[]User
		MockBannedUserList       *[]User
		MockEnableUser           *Response
		MockReactivateBannedUser *Response
		MockDisableUser          *Response
		MockDisableUserForCause  *Response
		MockConfig               *Config
		MockUpdateConfigRes      *Response
	}
	type args struct {
		email string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *User
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockUser: &User{
					ID: 5,
				},
			},
			args: args{
				email: "test",
			},
			want: &User{
				ID: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MockDelegate{
				MockAddUserResp:          tt.fields.MockAddUserResp,
				MockUpdateUserResp:       tt.fields.MockUpdateUserResp,
				MockRole:                 tt.fields.MockRole,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockActivateCommentRes:   tt.fields.MockActivateCommentRes,
				MockDeActivateCommentRes: tt.fields.MockDeActivateCommentRes,
				MockUnactivatedUserList:  tt.fields.MockUnactivatedUserList,
				MockBannedUserList:       tt.fields.MockBannedUserList,
				MockEnableUser:           tt.fields.MockEnableUser,
				MockReactivateBannedUser: tt.fields.MockReactivateBannedUser,
				MockDisableUser:          tt.fields.MockDisableUser,
				MockDisableUserForCause:  tt.fields.MockDisableUserForCause,
				MockConfig:               tt.fields.MockConfig,
				MockUpdateConfigRes:      tt.fields.MockUpdateConfigRes,
			}
			if got := d.GetUser(tt.args.email); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_GetUserByID(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddUserAuthResp      *ResponseID
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockAddLikeRes           *ResponseID
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockActivateCommentRes   *Response
		MockDeActivateCommentRes *Response
		MockUnactivatedUserList  *[]User
		MockBannedUserList       *[]User
		MockEnableUser           *Response
		MockReactivateBannedUser *Response
		MockDisableUser          *Response
		MockDisableUserForCause  *Response
		MockConfig               *Config
		MockUpdateConfigRes      *Response
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *User
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockUser: &User{
					ID: 5,
				},
			},
			args: args{
				id: 5,
			},
			want: &User{
				ID: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MockDelegate{
				MockAddUserResp:          tt.fields.MockAddUserResp,
				MockUpdateUserResp:       tt.fields.MockUpdateUserResp,
				MockRole:                 tt.fields.MockRole,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockActivateCommentRes:   tt.fields.MockActivateCommentRes,
				MockDeActivateCommentRes: tt.fields.MockDeActivateCommentRes,
				MockUnactivatedUserList:  tt.fields.MockUnactivatedUserList,
				MockBannedUserList:       tt.fields.MockBannedUserList,
				MockEnableUser:           tt.fields.MockEnableUser,
				MockReactivateBannedUser: tt.fields.MockReactivateBannedUser,
				MockDisableUser:          tt.fields.MockDisableUser,
				MockDisableUserForCause:  tt.fields.MockDisableUserForCause,
				MockConfig:               tt.fields.MockConfig,
				MockUpdateConfigRes:      tt.fields.MockUpdateConfigRes,
			}
			if got := d.GetUserByID(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_GetUserList(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddUserAuthResp      *ResponseID
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockAddLikeRes           *ResponseID
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockActivateCommentRes   *Response
		MockDeActivateCommentRes *Response
		MockUserList             *[]User
		MockUnactivatedUserList  *[]User
		MockBannedUserList       *[]User
		MockEnableUser           *Response
		MockReactivateBannedUser *Response
		MockDisableUser          *Response
		MockDisableUserForCause  *Response
		MockConfig               *Config
		MockUpdateConfigRes      *Response
	}
	tests := []struct {
		name   string
		fields fields
		want   *[]User
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockUserList: &[]User{
					{ID: 5},
					{ID: 6},
				},
			},

			want: &[]User{
				{ID: 5},
				{ID: 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MockDelegate{
				MockAddUserResp:          tt.fields.MockAddUserResp,
				MockUpdateUserResp:       tt.fields.MockUpdateUserResp,
				MockRole:                 tt.fields.MockRole,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockActivateCommentRes:   tt.fields.MockActivateCommentRes,
				MockDeActivateCommentRes: tt.fields.MockDeActivateCommentRes,
				MockUserList:             tt.fields.MockUserList,
				MockUnactivatedUserList:  tt.fields.MockUnactivatedUserList,
				MockBannedUserList:       tt.fields.MockBannedUserList,
				MockEnableUser:           tt.fields.MockEnableUser,
				MockReactivateBannedUser: tt.fields.MockReactivateBannedUser,
				MockDisableUser:          tt.fields.MockDisableUser,
				MockDisableUserForCause:  tt.fields.MockDisableUserForCause,
				MockConfig:               tt.fields.MockConfig,
				MockUpdateConfigRes:      tt.fields.MockUpdateConfigRes,
			}
			if got := d.GetUserList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.GetUserList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_GetUnActivatedUserList(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddUserAuthResp      *ResponseID
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockAddLikeRes           *ResponseID
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockActivateCommentRes   *Response
		MockDeActivateCommentRes *Response
		MockUserList             *[]User
		MockUnactivatedUserList  *[]User
		MockBannedUserList       *[]User
		MockEnableUser           *Response
		MockReactivateBannedUser *Response
		MockDisableUser          *Response
		MockDisableUserForCause  *Response
		MockConfig               *Config
		MockUpdateConfigRes      *Response
	}
	tests := []struct {
		name   string
		fields fields
		want   *[]User
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockUnactivatedUserList: &[]User{
					{ID: 5},
					{ID: 6},
				},
			},

			want: &[]User{
				{ID: 5},
				{ID: 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MockDelegate{
				MockAddUserResp:          tt.fields.MockAddUserResp,
				MockUpdateUserResp:       tt.fields.MockUpdateUserResp,
				MockRole:                 tt.fields.MockRole,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockActivateCommentRes:   tt.fields.MockActivateCommentRes,
				MockDeActivateCommentRes: tt.fields.MockDeActivateCommentRes,
				MockUserList:             tt.fields.MockUserList,
				MockUnactivatedUserList:  tt.fields.MockUnactivatedUserList,
				MockBannedUserList:       tt.fields.MockBannedUserList,
				MockEnableUser:           tt.fields.MockEnableUser,
				MockReactivateBannedUser: tt.fields.MockReactivateBannedUser,
				MockDisableUser:          tt.fields.MockDisableUser,
				MockDisableUserForCause:  tt.fields.MockDisableUserForCause,
				MockConfig:               tt.fields.MockConfig,
				MockUpdateConfigRes:      tt.fields.MockUpdateConfigRes,
			}
			if got := d.GetUnActivatedUserList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.GetUnActivatedUserList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_GetBannedUserList(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddUserAuthResp      *ResponseID
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockAddLikeRes           *ResponseID
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockActivateCommentRes   *Response
		MockDeActivateCommentRes *Response
		MockUserList             *[]User
		MockUnactivatedUserList  *[]User
		MockBannedUserList       *[]User
		MockEnableUser           *Response
		MockReactivateBannedUser *Response
		MockDisableUser          *Response
		MockDisableUserForCause  *Response
		MockConfig               *Config
		MockUpdateConfigRes      *Response
	}
	tests := []struct {
		name   string
		fields fields
		want   *[]User
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockBannedUserList: &[]User{
					{ID: 5},
					{ID: 6},
				},
			},

			want: &[]User{
				{ID: 5},
				{ID: 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MockDelegate{
				MockAddUserResp:          tt.fields.MockAddUserResp,
				MockUpdateUserResp:       tt.fields.MockUpdateUserResp,
				MockRole:                 tt.fields.MockRole,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockActivateCommentRes:   tt.fields.MockActivateCommentRes,
				MockDeActivateCommentRes: tt.fields.MockDeActivateCommentRes,
				MockUserList:             tt.fields.MockUserList,
				MockUnactivatedUserList:  tt.fields.MockUnactivatedUserList,
				MockBannedUserList:       tt.fields.MockBannedUserList,
				MockEnableUser:           tt.fields.MockEnableUser,
				MockReactivateBannedUser: tt.fields.MockReactivateBannedUser,
				MockDisableUser:          tt.fields.MockDisableUser,
				MockDisableUserForCause:  tt.fields.MockDisableUserForCause,
				MockConfig:               tt.fields.MockConfig,
				MockUpdateConfigRes:      tt.fields.MockUpdateConfigRes,
			}
			if got := d.GetBannedUserList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.GetBannedUserList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_EnableUser(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddUserAuthResp      *ResponseID
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockAddLikeRes           *ResponseID
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockActivateCommentRes   *Response
		MockDeActivateCommentRes *Response
		MockUserList             *[]User
		MockUnactivatedUserList  *[]User
		MockBannedUserList       *[]User
		MockEnableUser           *Response
		MockReactivateBannedUser *Response
		MockDisableUser          *Response
		MockDisableUserForCause  *Response
		MockConfig               *Config
		MockUpdateConfigRes      *Response
	}
	type args struct {
		u *User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockEnableUser: &Response{
					Success: true,
				},
			},
			args: args{
				u: &User{},
			},
			want: &Response{
				Success: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MockDelegate{
				MockAddUserResp:          tt.fields.MockAddUserResp,
				MockUpdateUserResp:       tt.fields.MockUpdateUserResp,
				MockRole:                 tt.fields.MockRole,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockActivateCommentRes:   tt.fields.MockActivateCommentRes,
				MockDeActivateCommentRes: tt.fields.MockDeActivateCommentRes,
				MockUserList:             tt.fields.MockUserList,
				MockUnactivatedUserList:  tt.fields.MockUnactivatedUserList,
				MockBannedUserList:       tt.fields.MockBannedUserList,
				MockEnableUser:           tt.fields.MockEnableUser,
				MockReactivateBannedUser: tt.fields.MockReactivateBannedUser,
				MockDisableUser:          tt.fields.MockDisableUser,
				MockDisableUserForCause:  tt.fields.MockDisableUserForCause,
				MockConfig:               tt.fields.MockConfig,
				MockUpdateConfigRes:      tt.fields.MockUpdateConfigRes,
			}
			if got := d.EnableUser(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.EnableUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_DisableUser(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddUserAuthResp      *ResponseID
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockAddLikeRes           *ResponseID
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockActivateCommentRes   *Response
		MockDeActivateCommentRes *Response
		MockUserList             *[]User
		MockUnactivatedUserList  *[]User
		MockBannedUserList       *[]User
		MockEnableUser           *Response
		MockReactivateBannedUser *Response
		MockDisableUser          *Response
		MockDisableUserForCause  *Response
		MockConfig               *Config
		MockUpdateConfigRes      *Response
	}
	type args struct {
		u *User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockDisableUser: &Response{
					Success: true,
				},
			},
			args: args{
				u: &User{},
			},
			want: &Response{
				Success: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MockDelegate{
				MockAddUserResp:          tt.fields.MockAddUserResp,
				MockUpdateUserResp:       tt.fields.MockUpdateUserResp,
				MockRole:                 tt.fields.MockRole,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockActivateCommentRes:   tt.fields.MockActivateCommentRes,
				MockDeActivateCommentRes: tt.fields.MockDeActivateCommentRes,
				MockUserList:             tt.fields.MockUserList,
				MockUnactivatedUserList:  tt.fields.MockUnactivatedUserList,
				MockBannedUserList:       tt.fields.MockBannedUserList,
				MockEnableUser:           tt.fields.MockEnableUser,
				MockReactivateBannedUser: tt.fields.MockReactivateBannedUser,
				MockDisableUser:          tt.fields.MockDisableUser,
				MockDisableUserForCause:  tt.fields.MockDisableUserForCause,
				MockConfig:               tt.fields.MockConfig,
				MockUpdateConfigRes:      tt.fields.MockUpdateConfigRes,
			}
			if got := d.DisableUser(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.DisableUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_DisableUserForCause(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddUserAuthResp      *ResponseID
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockAddLikeRes           *ResponseID
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockActivateCommentRes   *Response
		MockDeActivateCommentRes *Response
		MockUserList             *[]User
		MockUnactivatedUserList  *[]User
		MockBannedUserList       *[]User
		MockEnableUser           *Response
		MockReactivateBannedUser *Response
		MockDisableUser          *Response
		MockDisableUserForCause  *Response
		MockConfig               *Config
		MockUpdateConfigRes      *Response
	}
	type args struct {
		u *User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockDisableUserForCause: &Response{
					Success: true,
				},
			},
			args: args{
				u: &User{},
			},
			want: &Response{
				Success: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MockDelegate{
				MockAddUserResp:          tt.fields.MockAddUserResp,
				MockUpdateUserResp:       tt.fields.MockUpdateUserResp,
				MockRole:                 tt.fields.MockRole,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockActivateCommentRes:   tt.fields.MockActivateCommentRes,
				MockDeActivateCommentRes: tt.fields.MockDeActivateCommentRes,
				MockUserList:             tt.fields.MockUserList,
				MockUnactivatedUserList:  tt.fields.MockUnactivatedUserList,
				MockBannedUserList:       tt.fields.MockBannedUserList,
				MockEnableUser:           tt.fields.MockEnableUser,
				MockReactivateBannedUser: tt.fields.MockReactivateBannedUser,
				MockDisableUser:          tt.fields.MockDisableUser,
				MockDisableUserForCause:  tt.fields.MockDisableUserForCause,
				MockConfig:               tt.fields.MockConfig,
				MockUpdateConfigRes:      tt.fields.MockUpdateConfigRes,
			}
			if got := d.DisableUserForCause(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.DisableUserForCause() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_ReactivateUser(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddUserAuthResp      *ResponseID
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockAddLikeRes           *ResponseID
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockActivateCommentRes   *Response
		MockDeActivateCommentRes *Response
		MockUserList             *[]User
		MockUnactivatedUserList  *[]User
		MockBannedUserList       *[]User
		MockEnableUser           *Response
		MockReactivateBannedUser *Response
		MockDisableUser          *Response
		MockDisableUserForCause  *Response
		MockConfig               *Config
		MockUpdateConfigRes      *Response
	}
	type args struct {
		u *User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockReactivateBannedUser: &Response{
					Success: true,
				},
			},
			args: args{
				u: &User{},
			},
			want: &Response{
				Success: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MockDelegate{
				MockAddUserResp:          tt.fields.MockAddUserResp,
				MockUpdateUserResp:       tt.fields.MockUpdateUserResp,
				MockRole:                 tt.fields.MockRole,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockActivateCommentRes:   tt.fields.MockActivateCommentRes,
				MockDeActivateCommentRes: tt.fields.MockDeActivateCommentRes,
				MockUserList:             tt.fields.MockUserList,
				MockUnactivatedUserList:  tt.fields.MockUnactivatedUserList,
				MockBannedUserList:       tt.fields.MockBannedUserList,
				MockEnableUser:           tt.fields.MockEnableUser,
				MockReactivateBannedUser: tt.fields.MockReactivateBannedUser,
				MockDisableUser:          tt.fields.MockDisableUser,
				MockDisableUserForCause:  tt.fields.MockDisableUserForCause,
				MockConfig:               tt.fields.MockConfig,
				MockUpdateConfigRes:      tt.fields.MockUpdateConfigRes,
			}
			if got := d.ReactivateUser(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.ReactivateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_AddRole(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockAddUserAuthResp      *ResponseID
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockAddLikeRes           *ResponseID
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockActivateCommentRes   *Response
		MockDeActivateCommentRes *Response
		MockUserList             *[]User
		MockUnactivatedUserList  *[]User
		MockBannedUserList       *[]User
		MockEnableUser           *Response
		MockReactivateBannedUser *Response
		MockDisableUser          *Response
		MockDisableUserForCause  *Response
		MockConfig               *Config
		MockUpdateConfigRes      *Response
	}
	type args struct {
		r *Role
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ResponseID
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockAddRoleResp: &ResponseID{
					ID:      4,
					Success: true,
				},
			},
			args: args{
				r: &Role{},
			},
			want: &ResponseID{
				ID:      4,
				Success: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MockDelegate{
				MockAddUserResp:          tt.fields.MockAddUserResp,
				MockUpdateUserResp:       tt.fields.MockUpdateUserResp,
				MockRole:                 tt.fields.MockRole,
				MockAddRoleResp:          tt.fields.MockAddRoleResp,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockActivateCommentRes:   tt.fields.MockActivateCommentRes,
				MockDeActivateCommentRes: tt.fields.MockDeActivateCommentRes,
				MockUserList:             tt.fields.MockUserList,
				MockUnactivatedUserList:  tt.fields.MockUnactivatedUserList,
				MockBannedUserList:       tt.fields.MockBannedUserList,
				MockEnableUser:           tt.fields.MockEnableUser,
				MockReactivateBannedUser: tt.fields.MockReactivateBannedUser,
				MockDisableUser:          tt.fields.MockDisableUser,
				MockDisableUserForCause:  tt.fields.MockDisableUserForCause,
				MockConfig:               tt.fields.MockConfig,
				MockUpdateConfigRes:      tt.fields.MockUpdateConfigRes,
			}
			if got := d.AddRole(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.AddRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_GetRole(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockAddUserAuthResp      *ResponseID
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockAddLikeRes           *ResponseID
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockActivateCommentRes   *Response
		MockDeActivateCommentRes *Response
		MockUserList             *[]User
		MockUnactivatedUserList  *[]User
		MockBannedUserList       *[]User
		MockEnableUser           *Response
		MockReactivateBannedUser *Response
		MockDisableUser          *Response
		MockDisableUserForCause  *Response
		MockConfig               *Config
		MockUpdateConfigRes      *Response
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Role
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockRole: &Role{
					ID: 5,
				},
			},
			args: args{
				name: "test",
			},
			want: &Role{
				ID: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MockDelegate{
				MockAddUserResp:          tt.fields.MockAddUserResp,
				MockUpdateUserResp:       tt.fields.MockUpdateUserResp,
				MockRole:                 tt.fields.MockRole,
				MockAddRoleResp:          tt.fields.MockAddRoleResp,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockActivateCommentRes:   tt.fields.MockActivateCommentRes,
				MockDeActivateCommentRes: tt.fields.MockDeActivateCommentRes,
				MockUserList:             tt.fields.MockUserList,
				MockUnactivatedUserList:  tt.fields.MockUnactivatedUserList,
				MockBannedUserList:       tt.fields.MockBannedUserList,
				MockEnableUser:           tt.fields.MockEnableUser,
				MockReactivateBannedUser: tt.fields.MockReactivateBannedUser,
				MockDisableUser:          tt.fields.MockDisableUser,
				MockDisableUserForCause:  tt.fields.MockDisableUserForCause,
				MockConfig:               tt.fields.MockConfig,
				MockUpdateConfigRes:      tt.fields.MockUpdateConfigRes,
			}
			if got := d.GetRole(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.GetRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_GetRoleList(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockAddUserAuthResp      *ResponseID
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockAddLikeRes           *ResponseID
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockActivateCommentRes   *Response
		MockDeActivateCommentRes *Response
		MockUserList             *[]User
		MockUnactivatedUserList  *[]User
		MockBannedUserList       *[]User
		MockEnableUser           *Response
		MockReactivateBannedUser *Response
		MockDisableUser          *Response
		MockDisableUserForCause  *Response
		MockConfig               *Config
		MockUpdateConfigRes      *Response
	}
	tests := []struct {
		name   string
		fields fields
		want   *[]Role
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockRoleList: &[]Role{
					{ID: 5},
					{ID: 6},
				},
			},

			want: &[]Role{
				{ID: 5},
				{ID: 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MockDelegate{
				MockAddUserResp:          tt.fields.MockAddUserResp,
				MockUpdateUserResp:       tt.fields.MockUpdateUserResp,
				MockRole:                 tt.fields.MockRole,
				MockAddRoleResp:          tt.fields.MockAddRoleResp,
				MockRoleList:             tt.fields.MockRoleList,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockActivateCommentRes:   tt.fields.MockActivateCommentRes,
				MockDeActivateCommentRes: tt.fields.MockDeActivateCommentRes,
				MockUserList:             tt.fields.MockUserList,
				MockUnactivatedUserList:  tt.fields.MockUnactivatedUserList,
				MockBannedUserList:       tt.fields.MockBannedUserList,
				MockEnableUser:           tt.fields.MockEnableUser,
				MockReactivateBannedUser: tt.fields.MockReactivateBannedUser,
				MockDisableUser:          tt.fields.MockDisableUser,
				MockDisableUserForCause:  tt.fields.MockDisableUserForCause,
				MockConfig:               tt.fields.MockConfig,
				MockUpdateConfigRes:      tt.fields.MockUpdateConfigRes,
			}
			if got := d.GetRoleList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.GetRoleList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_DeleteRole(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockAddLikeRes           *ResponseID
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockActivateCommentRes   *Response
		MockDeActivateCommentRes *Response
		MockUserList             *[]User
		MockUnactivatedUserList  *[]User
		MockBannedUserList       *[]User
		MockEnableUser           *Response
		MockReactivateBannedUser *Response
		MockDisableUser          *Response
		MockDisableUserForCause  *Response
		MockConfig               *Config
		MockUpdateConfigRes      *Response
	}
	type args struct {
		rid int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockDeleteRoleResp: &Response{
					Success: true,
				},
			},
			args: args{
				rid: 5,
			},
			want: &Response{
				Success: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MockDelegate{
				MockAddUserResp:          tt.fields.MockAddUserResp,
				MockUpdateUserResp:       tt.fields.MockUpdateUserResp,
				MockRole:                 tt.fields.MockRole,
				MockAddRoleResp:          tt.fields.MockAddRoleResp,
				MockRoleList:             tt.fields.MockRoleList,
				MockDeleteRoleResp:       tt.fields.MockDeleteRoleResp,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockActivateCommentRes:   tt.fields.MockActivateCommentRes,
				MockDeActivateCommentRes: tt.fields.MockDeActivateCommentRes,
				MockUserList:             tt.fields.MockUserList,
				MockUnactivatedUserList:  tt.fields.MockUnactivatedUserList,
				MockBannedUserList:       tt.fields.MockBannedUserList,
				MockEnableUser:           tt.fields.MockEnableUser,
				MockReactivateBannedUser: tt.fields.MockReactivateBannedUser,
				MockDisableUser:          tt.fields.MockDisableUser,
				MockDisableUserForCause:  tt.fields.MockDisableUserForCause,
				MockConfig:               tt.fields.MockConfig,
				MockUpdateConfigRes:      tt.fields.MockUpdateConfigRes,
			}
			if got := d.DeleteRole(tt.args.rid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.DeleteRole() = %v, want %v", got, tt.want)
			}
		})
	}
}
