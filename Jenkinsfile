pipeline {
  agent any
  stages {
    stage('Build Go') {
      steps {
        sh 'go build server.go'
      }
    }
  }
}