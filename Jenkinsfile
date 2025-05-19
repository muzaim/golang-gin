pipeline {
    agent {
        docker { image 'golang:1.20' }
    }

    stages {
        stage('Checkout') {
            steps {
                git url: 'https://github.com/muzaim/golang-gin', branch: 'main'
            }
        }

        stage('Install Dependencies') {
            steps {
                sh 'go mod tidy'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o crud-app'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Archive') {
            steps {
                archiveArtifacts artifacts: 'crud-app', onlyIfSuccessful: true
            }
        }
    }
}
