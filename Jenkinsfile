pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                git url: 'https://github.com/muzaim/golang-gin.git', branch: 'main'
            }
        }

        stage('Build inside Docker') {
            steps {
                sh '''
                docker run --rm -v $PWD:/app -w /app golang:1.20 sh -c "
                    go mod tidy &&
                    go build -o crud-app &&
                    go test ./...
                "
                '''
            }
        }
    }

    post {
        always {
            archiveArtifacts artifacts: 'crud-app', onlyIfSuccessful: true
        }
    }
}
