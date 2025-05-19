pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                // Clone dari GitHub
                git url: 'https://github.com/muzaim/golang-gin/new/', branch: 'main'
            }
        }

        stage('Build') {
            steps {
                echo 'Sedang membangun proyek...'
                // Contoh perintah shell
                sh 'date > output.txt'
                sh 'cat output.txt'
            }
        }

        stage('Archive') {
            steps {
                archiveArtifacts artifacts: 'output.txt', onlyIfSuccessful: true
            }
        }
    }
}
