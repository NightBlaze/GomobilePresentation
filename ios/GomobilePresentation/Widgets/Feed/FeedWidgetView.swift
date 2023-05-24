//
//  FeedWidgetView.swift
//  GomobilePresentation
//
//  Created by Alexander Timonenkov on 23.05.2023.
//

import AutoLayoutSugar
import Engine
import SwiftProtobuf
import UIKit

final class FeedWidgetView: BaseWidgetView,
                            FeedwidgetDisplayProtocol,
                            UICollectionViewDataSource,
                            UICollectionViewDelegate,
                            UICollectionViewDelegateFlowLayout {
    private lazy var widget = FeedwidgetCreate(self)
    private lazy var router = FeedRouter(widget: self)

    private var cellWidgets = Set<FeedItemCollectionViewCell>()

    private lazy var collectionView: UICollectionView = {
        let layout = UICollectionViewFlowLayout()
        layout.scrollDirection = .vertical

        let view = UICollectionView(frame: .zero, collectionViewLayout: layout).prepareForAutoLayout()
        view.dataSource = self
        view.delegate = self
        view.register(
            FeedItemCollectionViewCell.self,
            forCellWithReuseIdentifier: String(describing: FeedItemCollectionViewCell.self)
        )
        return view
    }()

    override init() {
        super.init()

        setupUI()
        setupLayout()
    }

    required init?(coder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }

    override func free() {
        cellWidgets.forEach { $0.free() }
        cellWidgets.removeAll()

        widget?.free()
        widget = nil
    }

    // MARK: - Overrides

    override func viewControllerDidLoad() {
        super.viewControllerDidLoad()

        setupInitialData()
    }

    override func viewControllerDidAppear(animated: Bool) {
        super.viewControllerDidAppear(animated: animated)

        widget?.fetchFeedAsync()
    }

    // MARK: - Private

    private func setupUI() {
        addSubview(collectionView)
    }

    private func setupLayout() {
        collectionView.pinToSuperview()
    }

    private func setupInitialData() {
    }

    private func visibleCells() -> [FeedItemCollectionViewCell] {
        collectionView
            .visibleCells
            .compactMap { $0 as? FeedItemCollectionViewCell }
    }

    // MARK: - FeedwidgetDisplayProtocol

    func localizationDidChange(_ data: Data?) {
        mainAsync { [weak self] in
            guard
                let self = self,
                let data = data,
                let viewModel = try? FeedLocalizationDidChangeViewModel(serializedData: data)
            else {
                return
            }

            self.visibleCells().forEach { cell in
                for feedItem in viewModel.feedItems where feedItem.id == cell.viewModel?.id {
                    cell.viewModel?.caption = feedItem.caption
                }
            }
        }
    }

    func feedDidFetch() {
        mainAsync { [weak self] in
            guard let self = self else { return }
            self.collectionView.reloadData()
        }
    }

    // MARK: - UICollectionViewDataSource

    func collectionView(_ collectionView: UICollectionView, numberOfItemsInSection section: Int) -> Int {
        widget?.feedItemsCount() ?? 0
    }

    func collectionView(
        _ collectionView: UICollectionView,
        cellForItemAt indexPath: IndexPath
    ) -> UICollectionViewCell {
        guard
            let data = widget?.feedItem(at: indexPath.row),
            let viewModel = try? FeedItemViewModel(serializedData: data),
            let cell = collectionView.dequeueReusableCell(
                withReuseIdentifier: String(describing: FeedItemCollectionViewCell.self),
                for: indexPath
            ) as? FeedItemCollectionViewCell
        else {
            return UICollectionViewCell()
        }

        cell.viewModel = viewModel
        cell.downloadImage()
        cellWidgets.insert(cell)
        return cell
    }

    // MARK: - UICollectionViewDelegate

    func collectionView(
        _ collectionView: UICollectionView,
        didEndDisplaying cell: UICollectionViewCell,
        forItemAt indexPath: IndexPath
    ) {
        guard let cell = cell as? FeedItemCollectionViewCell else { return }
        cell.reset()
    }

    // MARK: - UICollectionViewDelegateFlowLayout

    func collectionView(
        _ collectionView: UICollectionView,
        layout collectionViewLayout: UICollectionViewLayout,
        sizeForItemAt indexPath: IndexPath
    ) -> CGSize {
        CGSize(width: collectionView.bounds.width, height: 86)
    }
}
