k8s-intro
https://labs.play-with-k8s.com/

Create a master node

1: kubeadm init --apiserver-advertise-address $(hostname -i) --pod-network-cidr 10.5.0.0/16

2: kubectl apply -f https://raw.githubusercontent.com/cloudnativelabs/kube-router/master/daemonset/kubeadm-kuberouter.yaml

3: optional kubectl apply -f https://raw.githubusercontent.com/kubernetes/website/master/content/en/examples/application/nginx-app.yaml

optional? create and join a worker node nodes: kubeadm join 192.168.0.18:6443 --token lctt91.pivqf3rgxw19b874
--discovery-token-ca-cert-hash sha256:71f2c3153a00d3ec7c8a4813b5fa37f77aa2829c21e41429d1c163e754dd3fd2
