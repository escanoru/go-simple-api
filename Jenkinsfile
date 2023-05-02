pipeline {
    agent { label 'ansible' }
    options {
        buildDiscarder(logRotator(numToKeepStr: '7'))
    }
    environment {
        DOCKER_CREDS = credentials("${params.DOCKER_CREDENTIALS}")
        REPO_NAME = "${params.REPO_NAME}"
        CONT_TAG = "${params.TAG}"
        REGISTRY = "${params.REGISTRY}"
        
    }
    parameters {
            string(name: 'REPO_NAME', defaultValue: 'go-simple-api', description: 'Name of the app to push to Dockerhub')
            string(name: 'DOCKER_CREDENTIALS', defaultValue: 'docker-push', description: 'Name of the docker username and passwoed credential on Jenkins')
            string(name: 'TAG', defaultValue: 'empty', description: 'Tag name to use for the container build')
            choice(name: 'REGISTRY', choices: ['docker.io', 'quay.io'], description: 'Registry to use')
    }
    stages {
        stage('Build') {
            steps {   
                //  `loginctl enable-linger $USER` needed to avoid https://github.com/containers/podman/issues/10738, usually we just need to run it the 1st time.
                sh '''
                loginctl enable-linger $USER
                podman build -t ${DOCKER_CREDS_USR}/${REPO_NAME}:${CONT_TAG} .
                '''
            }
        }
        stage('Login') {
            steps {
                sh 'echo ${DOCKER_CREDS_PSW} | podman login ${REGISTRY} -u ${DOCKER_CREDS_USR} --password-stdin'
            }
        }
        stage('Push') {
            steps {
                // We can also used `podman push ${DOCKER_CREDS_USR}/${REPO_NAME}:${CONT_TAG} --creds ${DOCKER_CREDS_USR}:${DOCKER_CREDS_PSW}'` instead of using the Login stage 
                sh 'podman push ${DOCKER_CREDS_USR}/${REPO_NAME}:${CONT_TAG}'
            }
        }
    }
    post {
        always {
            sh "podman logout ${params.REGISTRY}"
        }
    }
}

