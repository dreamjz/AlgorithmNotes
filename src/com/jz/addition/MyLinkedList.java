package com.jz.addition;

import java.util.logging.Logger;

/**
 * �����ʵ��
 * 
 * @author Kesa
 *
 */
public class MyLinkedList {
	Node headNode;// ͷ���
	Logger log = Logger.getLogger("MyLinkedList");

	/**
	 * �ڵ���
	 * 
	 * @author Kesa
	 *
	 */
	class Node {
		int data;// �ڵ�����
		Node nextNode;// ָ����һ�ڵ������

		public Node(int data) {
			this.data = data;
		}
	}

	/**
	 * �����ڵ�
	 * 
	 * @param data �ڵ�����
	 */
	public void addNode(int data) {
		Node newNode = new Node(data);
		Node tmpNode = headNode;
		if (tmpNode == null) {// ����ͷ���
			headNode = newNode;
		} else {// ��������ĩβ
			while (tmpNode.nextNode != null) {
				tmpNode = tmpNode.nextNode;
			}
			tmpNode.nextNode = newNode;
		}
	}

	/**
	 * ɾ��ָ��λ�õĽڵ�
	 * 
	 * @param index ָ������
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
	 * ��ȡ������
	 * 
	 * @return ������
	 */
	public int getLength() {
		int count = 0;
		while (headNode.nextNode != null) {
			count++;
		}
		return count;
	}

	/**
	 * ɾ��ָ�����ݵĽڵ�
	 * 
	 * @param node ָ���ڵ�
	 * @return true ɾ���ɹ� false δ�ҵ��ڵ�
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
	 * �������
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
