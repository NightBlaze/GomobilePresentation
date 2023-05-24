//
//  BaseWidgetView.swift
//  GomobilePresentation
//
//  Created by Alexander Timonenkov on 23.05.2023.
//

import UIKit

class BaseWidgetView: UIView {
    private weak var _hostViewController: BaseViewController?
    var hostViewController: BaseViewController? {
        guard _hostViewController == nil else { return _hostViewController }

        return (superview as? BaseWidgetView)?.hostViewController
    }

    init() {
        super.init(frame: .zero)
    }

    required init?(coder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }

    func free() {
        assert(false, "should be overriden in subclass")
    }

    // MARK: - UIViewController

    @discardableResult
    func setHostViewController(_ hostViewController: BaseViewController?) -> Self {
        self._hostViewController = hostViewController
        hostViewController?.addWidgetView(self)
        return self
    }

    func viewControllerDidLoad() { }

    func viewControllerDidLayoutSubviews() { }

    func viewControllerWillAppear(animated: Bool) { }

    func viewControllerDidAppear(animated: Bool) { }

    func viewControllerWillDisappear(animated: Bool) { }

    func viewControllerDidDisappear(animated: Bool) { }

    func viewControllerWillTransition(to size: CGSize, with coordinator: UIViewControllerTransitionCoordinator) { }
}
