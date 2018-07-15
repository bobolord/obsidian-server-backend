pipeline {
  agent any
  stages {
    stage('Build Go') {
      steps {
        sh '''cd ~/go/src/github.com/bobolord/obsidian-server-backend
go build server.go'''
      }
    }
  }
}