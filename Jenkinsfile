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
                sh 'docker build --tag sorter .'
            }
        }
        stage('Upload to ECR') {
            steps {
                sh 'aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 921884257724.dkr.ecr.us-east-2.amazonaws.com/app-repository'
                sh 'docker tag sorter 921884257724.dkr.ecr.us-east-2.amazonaws.com/app-repository:1.1'
                sh 'docker push 921884257724.dkr.ecr.us-east-2.amazonaws.com/app-repository:1.1'
            }
        }
        stage('Update Service') {
            steps {
                script{
                    def task-definition = sh(returnStdout: true, script: 'aws ecs describe-task-definition --task-definition sorter-app')
                }
                sh 'aws ecs register-task-definition --region us-east-2 --family sorter-service --container-definitions ${task-definition}'
                sh 'aws ecs update-service --region us-east-2 --cluster web-app-cluster --service sorter-service --force-new-deployment'
            }
        }
    }
}