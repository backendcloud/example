import java.text.SimpleDateFormat
dateformat = new SimpleDateFormat("yyyyMMddHHmmss")
def time = dateformat.format(new Date())
def label = "jenkins-slave-${UUID.randomUUID().toString()}"
def images = []
def namespace = "49099d38-6b42-4f98-b872-cc91c87b16f4"
def scmproject = "kubevirt"
def buildNum = BUILD_NUMBER
def pipelineName = "KubeVirt-V0.53"
podTemplate(label: label, cloud:"kubernetes", nodeSelector:"jenkins-slave=true",containers: [
        containerTemplate(
                name: 'jnlp',
                image: 'registry.paas/cmss/jnlp-slave:3.10-1-alpine',
                args: '${computer.jnlpmac} ${computer.name}'),
        containerTemplate(
                name: 'golang-1-18',
                image: 'registry.paas/cmss/golang:1.18.5',
                command: 'cat',
                ttyEnabled: true),
        containerTemplate(
                name: 'harbor',
                image: 'registry.paas/cmss/harbor:v0.1',
                command: 'cat',
                ttyEnabled: true),
        containerTemplate(
                name: 'gotest',
                image: 'registry.paas/cmss/go-test:v0.1',
                command: 'cat',
                ttyEnabled: true),
        containerTemplate(
                name: 'sonar',
                image: 'registry.paas/cmss/sonar:v0.5',
                command: 'cat',
                ttyEnabled: true),
        containerTemplate(
                name: 'centos',
                image: 'docker.io/centos:7.9.2009',
                command: 'cat',
                ttyEnabled: true),

],
        volumes: [
                hostPathVolume(mountPath: '/var/run/docker.sock', hostPath: '/var/run/docker.sock'),
                hostPathVolume(mountPath: '/root/.m2/', hostPath: '/dcos/maven/.m2/'),
                hostPathVolume(mountPath: '/etc/hosts', hostPath: '/etc/hosts'),
                hostPathVolume(mountPath: '/root/.gradle', hostPath: '/dcos/gradle/.gradle')
        ]) {
    node(label) {
        stage("CodePull"){
            git branch: '${BRANCH}',credentialsId: 'gerrit-code-clone', url: 'http://gerrit.cmss.com/a/EKI/kubevirt'
            env.shortCommit = sh(returnStdout: true, script: "git log -n 1 --pretty=format:'%h'").trim()
            env.gitUrl = sh(returnStdout: true, script: 'git config remote.origin.url').trim()
            env.branch = sh(returnStdout: true, script: "git rev-parse --abbrev-ref HEAD").trim()
            env.commitMessage = sh(returnStdout: true, script: "git log --oneline --format=%B -n 1 HEAD").trim()
            sh """
                           mkdir ${namespace}-${scmproject}
                           mv `ls -1 | grep -v ${namespace}-${scmproject}` ${namespace}-${scmproject}
                           mv ${namespace}-${scmproject} ${scmproject}
                           ls -l
                           pwd
uname -a
apt-get install docker
                           """
        }
        stage("CodeBuild"){
            container('centos'){
                sh """
                                    cd ${scmproject}
                                    pwd
ls -l /var/run/docker.sock
                                    yum install -y make
                                    make
                                 """
            }

        }





    }}