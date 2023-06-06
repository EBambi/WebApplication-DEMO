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
/*         stage('Unit Testing') {
            steps 'go test'
        }
 */     stage('Build') {
            steps {
                script{
                    app = docker.build("sorter-app-image")
                }
            }
        }
        stage('Upload to ECR') {
            steps {
                aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 921884257724.dkr.ecr.us-east-2.amazonaws.com/app-repository
            }
        } 
    }
}