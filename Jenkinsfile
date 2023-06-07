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
            sh 'go test'
        }
        stage('Build') {
            steps {
                sh 'docker build --tag sorter .'
            }
        }
        stage('Upload to ECR') {
            steps {
                sh 'aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 921884257724.dkr.ecr.us-east-2.amazonaws.com/app-repository'
                sh 'docker tag sorter 921884257724.dkr.ecr.us-east-2.amazonaws.com/app-repository:latest'
                sh 'docker push 921884257724.dkr.ecr.us-east-2.amazonaws.com/app-repository:latest'
            }
        }
        stage('Update Service') {
            steps {
                sh 'aws ecs update-service --region us-east-2 --cluster web-app-cluster --service sorter-service --force-new-deployment'
            }
        }
    }
}