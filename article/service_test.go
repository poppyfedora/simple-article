package article

import (
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestServiceImpl_GetLists(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockRepo := NewMockRepository(mockCtrl)

	type fields struct {
		repo Repository
	}
	type args struct {
		keyword string
		author  string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     []Article
		wantErr  bool
		mockFunc func()
	}{
		{
			name:    "valid case",
			fields:  fields{repo: mockRepo},
			want:    []Article{},
			wantErr: false,
			mockFunc: func() {
				mockRepo.EXPECT().GetList(gomock.Any(), gomock.Any()).Return([]Article{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ServiceImpl{
				repo: tt.fields.repo,
			}
			tt.mockFunc()
			got, err := s.GetLists(tt.args.keyword, tt.args.author)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.GetLists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServiceImpl.GetLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceImpl_AddNew(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockRepo := NewMockRepository(mockCtrl)

	type fields struct {
		repo Repository
	}
	type args struct {
		article Article
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		mockFunc func()
	}{
		{
			name:    "valid case",
			fields:  fields{repo: mockRepo},
			args:    args{article: Article{}},
			wantErr: false,
			mockFunc: func() {
				mockRepo.EXPECT().Add(gomock.Any()).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ServiceImpl{
				repo: tt.fields.repo,
			}
			tt.mockFunc()
			if err := s.AddNew(tt.args.article); (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.AddNew() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
