pipeline {
    agent any

    environment {
        GO111MODULE = 'on'
    }

    stages {
        stage('Checkout') {
            steps {
                git url: 'https://github.com/muzaim/golang-gin', branch: 'main'
            }
        }

        stage('Install Dependencies') {
            steps {
                echo 'Installing Go modules...'
                sh 'go mod tidy'
            }
        }

        stage('Build') {
            steps {
                echo 'Building Go project...'
                sh 'go build -o crud-app'
            }
        }

        stage('Test') {
            steps {
                echo 'Running tests...'
                sh 'go test ./...'
            }
        }

        stage('Archive') {
            steps {
                archiveArtifacts artifacts: 'crud-app', onlyIfSuccessful: true
            }
        }
    }

    post {
        success {
            echo 'Build successful!'
        }
        failure {
            echo 'Build failed!'
        }
    }
}
