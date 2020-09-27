<div align=center>

# Dev Event Server

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)]

</div>

<br />

## Tech

- Go
- Gin-gonic
- JWT Auth
- MySQL 
- Go-queryset
- Docker
- swagger
- Cloud
  - AWS ECB or EC2 or NCloud...
- CI/CD
  - Test code
  - Travis CI or Circle CI 
  - Awesome GO
  - deepsource

## 기능

- 인증
  - [x] 로그인
    - [x] JWT 토큰 발급
  - [ ] 로그아웃 - Client
  - [x] 회원가입
- 프로필
  - [x] 사용자 정보 불러오기
  - [x] 사용자 정보 수정
- 관리자
  - 이벤트
    - [ ] 이벤트 전체보기
    - [ ] 이벤트 추가
    - [ ] 이벤트 수정
    - [x] 이벤트 삭제
  - 태그 
    - [x] 태그 전체보기
    - [x] 태그 생성 
    - [x] 태그 수정
    - [x] 태그 삭제
  - 사용자
    - [x] 회원 정보들 가져오기
    - [x] 회원 정보 열람
    - [x] 회원 정보 수정
      - [ ] 비밀번호 변경
    - [x] 회원 정보 삭제
- 이벤트
  - [ ] 이벤트 현재 달 기준 조회
  - [ ] 이벤트 태그로 조회



# Login

```
{
    "email": "apple12",
    "password": "1234"
}
```

# API 

[ 이벤트 확인 페이지 ] 

```text
현재 달, 월의 이벤트 확인
/events

달, 월로 확인
/events/:year/:month
```

<br />


[ 관리자  - 이벤트 ] 

```text
이벤트 추가 
  @Auth. 관리자 
[POST] /admin/events

이벤트 수정
  @Auth. 관리자 
[PUT] /admin/events/:eventID

이벤트 정보 가져오기
  @Auth. 관리자 
[GET] /admin/event/:eventID
```

<br />

[ 관리자  - 회원 정보 ] 

```text
[DONE] 회원 정보들 가져오기 
  @Auth. 관리자 
[GET] /admin/users

[DONE] 회원 정보 열람
  @Auth. 관리자 
[GET] /admin/users/:userID

[DONE] 회원 정보 수정
  @Auth. 관리자 
[PUT] /admin/users/:userID

[DONE] 회원 삭제
  @Auth. 관리자 
[DELETE] /admin/users/:userID
```

<br />

-----------------------------

<br />

[2차 개발] 자유게시판 

```text
전체 글 목록 보기 
  @Param: Offset
  @Param: Limit
[GET] /free-board

글 읽기
[GET] /free-board/:postID

글 삭제
  @Auth: 글쓴이 인증
[DELETE] /free-board/:postID

글 작성
  @Auth: 로그인 사용자
[Post] /free-board/:postID

글 수정
  @Auth: 글쓴이 인증
[Put] /free-board/:postID

덧글작성 
  @Auth: 로그인 사용자
[Post] /free-board/:postID/comments

덧글 수정
  @Auth: 덧글 작성자
[Put] /free-board/:postID/comments/:commentID

덧글 삭제
  @Auth: 덧글 작성자
[Delete] /free-board/:postID/comments/:commentID
```