package com.jz.addition;

import java.util.logging.Logger;

/**
 * 链表的实现
 * 
 * @author Kesa
 *
 */
public class MyLinkedList {
	Node headNode;// 头结点
	Logger log = Logger.getLogger("MyLinkedList");

	/**
	 * 节点类
	 * 
	 * @author Kesa
	 *
	 */
	class Node {
		int data;// 节点数据
		Node nextNode;// 指向下一节点的引用

		public Node(int data) {
			this.data = data;
		}
	}

	/**
	 * 新增节点
	 * 
	 * @param data 节点数据
	 */
	public void addNode(int data) {
		Node newNode = new Node(data);
		Node tmpNode = headNode;
		if (tmpNode == null) {// 设置头结点
			headNode = newNode;
		} else {// 插至链表末尾
			while (tmpNode.nextNode != null) {
				tmpNode = tmpNode.nextNode;
			}
			tmpNode.nextNode = newNode;
		}
	}

	/**
	 * 删除指定位置的节点
	 * 
	 * @param index 指定索引
	 * @return
	 */
	public boolean deleteNode(int index) {
		if (index < 0 || index > getLength()) {
			return false;
		}
		int i = 0;
		Node tmpNode = headNode;
		Node preNode = tmpNode;
		while (i != index) {
			preNode = tmpNode;
			tmpNode = tmpNode.nextNode;
			i++;
		}
		preNode.nextNode = tmpNode.nextNode;
		return true;
	}

	/**
	 * 获取链表长度
	 * 
	 * @return 链表长度
	 */
	public int getLength() {
		int count = 0;
		while (headNode.nextNode != null) {
			count++;
		}
		return count;
	}

	/**
	 * 删除指定内容的节点
	 * 
	 * @param node 指定节点
	 * @return true 删除成功 false 未找到节点
	 */
	public boolean deleteNode(Node node) {
		Node tmpNode = headNode;
		Node preNode = tmpNode;
		while ((tmpNode.data != node.data) || (tmpNode.nextNode != null)) {
			preNode = tmpNode;
			tmpNode = tmpNode.nextNode;
		}
		if (tmpNode.data == node.data) {
			preNode.nextNode = tmpNode.nextNode;
			return true;
		} else {
			return false;
		}
	}

	/**
	 * 输出链表
	 */
	public void printLinkedList() {
		StringBuilder sb = new StringBuilder("[");
		Node tmpNode = headNode;
		while (tmpNode.nextNode != null) {
			sb.append(tmpNode.data);
			sb.append(",");
			tmpNode = tmpNode.nextNode;
		}
		sb.append(tmpNode.data);
		sb.append("]");
	}

}
