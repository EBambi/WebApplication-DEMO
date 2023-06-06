pipeline {
    agent {label 'agent-one'}
    stages{
        stage('Checkout') {
            steps {
                cleanWs()
                sshagent(credentials: ['WebApplication-Credentials']) {
                    git url: 'https://github.com/EBambi/WebApplication-DEMO.git', branch: 'master', credentialsId: 'WebApplication-Credentials'
                }
            }
        }
        stage('Unit Testing') {
            steps 'go test'
        }
/*        stage('Build') {
            steps {
                script{
                    app = docker.build("sorter-app-image")
                }
            }
        }
        stage('Upload to ECR') {
            steps {
                docker.withRegistry('https://921884257724.dkr.ecr.us-east-2.amazonaws.com/app-repository'){
                    app.push("${env.BUILD_NUMBER}")
                    app.push("latest")
                }
            }
        } */
    }
}