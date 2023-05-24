//
//  BaseViewController.swift
//  GomobilePresentation
//
//  Created by Alexander Timonenkov on 23.05.2023.
//

import UIKit

class BaseViewController: UIViewController {
    private var widgetViews = NSHashTable<BaseWidgetView>(options: .weakMemory)

    // MARK: - View life cycle

    override func viewDidLoad() {
        super.viewDidLoad()

        widgetViews.allObjects.forEach { $0.viewControllerDidLoad() }
    }

    override func viewDidLayoutSubviews() {
        super.viewDidLayoutSubviews()

        widgetViews.allObjects.forEach { $0.viewControllerDidLayoutSubviews() }
    }

    override func viewWillAppear(_ animated: Bool) {
        super.viewWillAppear(animated)

        widgetViews.allObjects.forEach { $0.viewControllerWillAppear(animated: animated) }
    }

    override func viewDidAppear(_ animated: Bool) {
        super.viewDidAppear(animated)

        widgetViews.allObjects.forEach { $0.viewControllerDidAppear(animated: animated) }
    }

    override func viewWillDisappear(_ animated: Bool) {
        super.viewWillDisappear(animated)

        widgetViews.allObjects.forEach { $0.viewControllerWillDisappear(animated: animated) }
    }

    override func viewDidDisappear(_ animated: Bool) {
        super.viewDidDisappear(animated)

        widgetViews.allObjects.forEach { $0.viewControllerDidDisappear(animated: animated) }
    }

    override func viewWillTransition(to size: CGSize, with coordinator: UIViewControllerTransitionCoordinator) {
        super.viewWillTransition(to: size, with: coordinator)

        widgetViews.allObjects.forEach { $0.viewControllerWillTransition(to: size, with: coordinator) }
    }

    // MARK: - Internal functions

    /// Should be called only from BaseWidgetView
    func addWidgetView(_ widgetView: BaseWidgetView) {
        widgetViews.add(widgetView)
    }
}

