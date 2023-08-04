package k8s

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
	"testing"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func onAdd(obj interface{}) {
	pod := obj.(*corev1.Pod)
	fmt.Println("add a pod:", pod.Name)
}

func TestK8sInformer(t *testing.T) {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	// 初始化 client
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panic(err.Error())
	}

	stopper := make(chan struct{})
	defer close(stopper)

	// 初始化 informer
	factory := informers.NewSharedInformerFactory(clientSet, 0)
	podInformer := factory.Core().V1().Pods()
	informer := podInformer.Informer()
	defer runtime.HandleCrash()

	// 启动 informer，list & watch
	go factory.Start(stopper)

	// 从 apiserver 同步资源，即 list
	if !cache.WaitForCacheSync(stopper, informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	// 使用自定义 handler
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAdd,
		UpdateFunc: func(interface{}, interface{}) { fmt.Println("update not implemented") }, // 此处省略 workqueue 的使用
		DeleteFunc: func(interface{}) { fmt.Println("delete not implemented") },
	})

	// 创建 lister
	podLister := podInformer.Lister()

	// 从 lister 中获取所有 items
	podList, err := podLister.List(labels.Everything())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("podlist: %#v", podList)
	<-stopper
}
