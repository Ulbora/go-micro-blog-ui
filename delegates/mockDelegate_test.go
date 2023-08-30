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

func TestMockDelegate_AddBlog(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		b *Blog
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
				MockAddBlogRes: &ResponseID{
					ID:      4,
					Success: true,
				},
			},
			args: args{
				b: &Blog{},
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
				MockRoleList:             tt.fields.MockRoleList,
				MockDeleteRoleResp:       tt.fields.MockDeleteRoleResp,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.AddBlog(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.AddBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_UpdateBlog(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		b *Blog
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
				MockUpdateBlogRes: &Response{
					Success: true,
				},
			},
			args: args{
				b: &Blog{},
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
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.UpdateBlog(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.UpdateBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_GetBlog(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		bid int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Blog
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockBlog: &Blog{
					ID: 5,
				},
			},
			args: args{
				bid: 3,
			},
			want: &Blog{
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
				MockRoleList:             tt.fields.MockRoleList,
				MockDeleteRoleResp:       tt.fields.MockDeleteRoleResp,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.GetBlog(tt.args.bid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.GetBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_GetBlogByName(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		name  string
		start int64
		end   int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]Blog
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockBlogList: &[]Blog{
					{ID: 5},
					{ID: 6},
				},
			},
			args: args{
				name:  "test",
				start: 0,
				end:   10,
			},

			want: &[]Blog{
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
				MockDeleteRoleResp:       tt.fields.MockDeleteRoleResp,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.GetBlogByName(tt.args.name, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.GetBlogByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_GetAdminBlogByName(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		name  string
		start int64
		end   int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]Blog
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockBlogList: &[]Blog{
					{ID: 5},
					{ID: 6},
				},
			},
			args: args{
				name:  "test",
				start: 0,
				end:   10,
			},

			want: &[]Blog{
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
				MockDeleteRoleResp:       tt.fields.MockDeleteRoleResp,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.GetAdminBlogByName(tt.args.name, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.GetBlogByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_GetBlogList(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		start int64
		end   int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]Blog
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockBlogList: &[]Blog{
					{ID: 5},
					{ID: 6},
				},
			},
			args: args{
				start: 0,
				end:   10,
			},

			want: &[]Blog{
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
				MockDeleteRoleResp:       tt.fields.MockDeleteRoleResp,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.GetBlogList(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.GetBlogList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_GetAdminBlogList(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		start int64
		end   int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]Blog
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockBlogList: &[]Blog{
					{ID: 5},
					{ID: 6},
				},
			},
			args: args{
				start: 0,
				end:   10,
			},

			want: &[]Blog{
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
				MockDeleteRoleResp:       tt.fields.MockDeleteRoleResp,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.GetAdminBlogList(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.GetAdminBlogList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_ActivateBlog(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		b *Blog
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
				MockActivateRes: &Response{
					Success: true,
				},
			},
			args: args{
				b: &Blog{},
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
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.ActivateBlog(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.ActivateBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_DeActivateBlog(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		b *Blog
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
				MockDeActivateRes: &Response{
					Success: true,
				},
			},
			args: args{
				b: &Blog{},
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
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.DeActivateBlog(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.DeActivateBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_AddLike(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		l *Like
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
				MockAddLikeRes: &ResponseID{
					ID:      4,
					Success: true,
				},
			},
			args: args{
				l: &Like{},
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
				MockRoleList:             tt.fields.MockRoleList,
				MockDeleteRoleResp:       tt.fields.MockDeleteRoleResp,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.AddLike(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.AddLike() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_RemoveLike(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		l *Like
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
				MockRemoveLikeResp: &Response{

					Success: true,
				},
			},
			args: args{
				l: &Like{},
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
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.RemoveLike(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.RemoveLike() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_ViewLikes(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		bid int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]Like
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockLikeList: &[]Like{
					{BlogID: 5},
					{BlogID: 6},
				},
			},
			args: args{

				bid: 5,
			},
			want: &[]Like{
				{BlogID: 5},
				{BlogID: 6},
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
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.ViewLikes(tt.args.bid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.ViewLikes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_AddComment(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		c *Comment
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
				MockAddCommentRes: &ResponseID{
					ID:      4,
					Success: true,
				},
			},
			args: args{
				c: &Comment{},
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
				MockRoleList:             tt.fields.MockRoleList,
				MockDeleteRoleResp:       tt.fields.MockDeleteRoleResp,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.AddComment(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.AddComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_UpdateComment(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		c *Comment
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
				MockUpdateCommentRes: &Response{
					Success: true,
				},
			},
			args: args{
				c: &Comment{},
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
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.UpdateComment(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.UpdateComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_GetCommentList(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		blogID int64
		start  int64
		end    int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]Comment
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockCommentList: &[]Comment{
					{BlogID: 5},
					{BlogID: 6},
				},
			},
			args: args{
				blogID: 5,
				start:  0,
				end:    10,
			},
			want: &[]Comment{
				{BlogID: 5},
				{BlogID: 6},
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
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.GetCommentList(tt.args.blogID, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.GetCommentList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_GetAdminCommentList(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		blogID int64
		start  int64
		end    int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]Comment
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockCommentList: &[]Comment{
					{BlogID: 5},
					{BlogID: 6},
				},
			},
			args: args{
				blogID: 5,
				start:  0,
				end:    10,
			},
			want: &[]Comment{
				{BlogID: 5},
				{BlogID: 6},
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
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.GetAdminCommentList(tt.args.blogID, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.GetAdminCommentList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_ActivateComment(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		c *Comment
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
				MockActivateCommentRes: &Response{
					Success: true,
				},
			},
			args: args{
				c: &Comment{},
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
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.ActivateComment(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.ActivateComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_DeActivateComment(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		c *Comment
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
				MockDeActivateCommentRes: &Response{
					Success: true,
				},
			},
			args: args{
				c: &Comment{},
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
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.DeActivateComment(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.DeActivateComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_AddUserAuth(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		a *UserAuth
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
				MockAddUserAuthResp: &ResponseID{
					ID:      5,
					Success: true,
				},
			},
			args: args{
				a: &UserAuth{},
			},
			want: &ResponseID{
				ID:      5,
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
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.AddUserAuth(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.AddUserAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_GetUserAuthList(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		userID int64
		start  int64
		end    int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]UserAuth
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockUserAuthList: &[]UserAuth{
					{ID: 5},
					{ID: 6},
				},
			},
			args: args{
				userID: 5,
				start:  0,
				end:    10,
			},
			want: &[]UserAuth{
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
				MockDeleteRoleResp:       tt.fields.MockDeleteRoleResp,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.GetUserAuthList(tt.args.userID, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.GetUserAuthList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_UpdateConfig(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		c *Config
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
				MockUpdateConfigRes: &Response{
					Success: true,
				},
			},
			args: args{
				c: &Config{},
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
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.UpdateConfig(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.UpdateConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_GetConfig(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
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
		want   *Config
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockConfig: &Config{
					ID: 5,
				},
			},

			want: &Config{
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
				MockRoleList:             tt.fields.MockRoleList,
				MockDeleteRoleResp:       tt.fields.MockDeleteRoleResp,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
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
			if got := d.GetConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_DeleteBlog(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockDeleteBlogRes        *Response
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
		bid int64
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
				MockDeleteBlogRes: &Response{
					Success: true,
				},
			},
			args: args{
				bid: 5,
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
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockDeleteBlogRes:        tt.fields.MockDeleteBlogRes,
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
			if got := d.DeleteBlog(tt.args.bid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.DeleteBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_AddRule(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockDeleteBlogRes        *Response
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
		MockAddRuleResp          *ResponseID
	}
	type args struct {
		r *Rule
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
				MockAddRuleResp: &ResponseID{
					ID:      4,
					Success: true,
				},
			},
			args: args{
				r: &Rule{},
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
				MockRoleList:             tt.fields.MockRoleList,
				MockDeleteRoleResp:       tt.fields.MockDeleteRoleResp,
				MockAddUserAuthResp:      tt.fields.MockAddUserAuthResp,
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockDeleteBlogRes:        tt.fields.MockDeleteBlogRes,
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
				MockAddRuleResp:          tt.fields.MockAddRuleResp,
			}
			if got := d.AddRule(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.AddRule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_UpdateRule(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockDeleteBlogRes        *Response
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
		MockAddRuleResp          *ResponseID
		MockUpdateRuleResp       *Response
	}
	type args struct {
		r *Rule
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
				MockUpdateRuleResp: &Response{
					Success: true,
				},
			},
			args: args{
				r: &Rule{},
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
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockDeleteBlogRes:        tt.fields.MockDeleteBlogRes,
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
				MockAddRuleResp:          tt.fields.MockAddRuleResp,
				MockUpdateRuleResp:       tt.fields.MockUpdateRuleResp,
			}
			if got := d.UpdateRule(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.UpdateRule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDelegate_GetRule(t *testing.T) {
	type fields struct {
		MockAddUserResp          *ResponseID
		MockUpdateUserResp       *Response
		MockRole                 *Role
		MockAddRoleResp          *ResponseID
		MockRoleList             *[]Role
		MockDeleteRoleResp       *Response
		MockAddUserAuthResp      *ResponseID
		MockUserAuthList         *[]UserAuth
		MockUser                 *User
		MockBlogList             *[]Blog
		MockBlog                 *Blog
		MockCommentList          *[]Comment
		MockLikeList             *[]Like
		MockAddBlogRes           *ResponseID
		MockUpdateBlogRes        *Response
		MockAddCommentRes        *ResponseID
		MockUpdateCommentRes     *Response
		MockAddLikeRes           *ResponseID
		MockRemoveLikeResp       *Response
		MockActivateRes          *Response
		MockDeActivateRes        *Response
		MockDeleteBlogRes        *Response
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
		MockAddRuleResp          *ResponseID
		MockUpdateRuleResp       *Response
		MockRule                 *Rule
	}
	tests := []struct {
		name   string
		fields fields
		want   *Rule
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockRule: &Rule{
					ID:      4,
					Content: "test",
				},
			},
			
			want: &Rule{
				ID:      4,
				Content: "test",
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
				MockUserAuthList:         tt.fields.MockUserAuthList,
				MockUser:                 tt.fields.MockUser,
				MockBlogList:             tt.fields.MockBlogList,
				MockBlog:                 tt.fields.MockBlog,
				MockCommentList:          tt.fields.MockCommentList,
				MockLikeList:             tt.fields.MockLikeList,
				MockAddBlogRes:           tt.fields.MockAddBlogRes,
				MockUpdateBlogRes:        tt.fields.MockUpdateBlogRes,
				MockAddCommentRes:        tt.fields.MockAddCommentRes,
				MockUpdateCommentRes:     tt.fields.MockUpdateCommentRes,
				MockAddLikeRes:           tt.fields.MockAddLikeRes,
				MockRemoveLikeResp:       tt.fields.MockRemoveLikeResp,
				MockActivateRes:          tt.fields.MockActivateRes,
				MockDeActivateRes:        tt.fields.MockDeActivateRes,
				MockDeleteBlogRes:        tt.fields.MockDeleteBlogRes,
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
				MockAddRuleResp:          tt.fields.MockAddRuleResp,
				MockUpdateRuleResp:       tt.fields.MockUpdateRuleResp,
				MockRule:                 tt.fields.MockRule,
			}
			if got := d.GetRule(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDelegate.GetRule() = %v, want %v", got, tt.want)
			}
		})
	}
}
